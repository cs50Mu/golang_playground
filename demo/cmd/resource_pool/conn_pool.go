package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

// Conn 连接应该实现 Close 方法
type Conn interface {
	Close() error
}

// poolConn wraps a `Conn` with a mutex
type poolConn struct {
	pool *Pool
	sync.Mutex
	ci        Conn // connection inner, the real connection
	createdAt time.Time
	closed    bool
}

func (pc *poolConn) Close() error {
	pc.Lock()
	if pc.closed {
		pc.Unlock()
		return errors.New("duplicate poolConn close")
	}
	pc.closed = true
	pc.Unlock()

	return pc.ci.Close()
}

// ConnFactory 连接生成
type ConnFactory func() (Conn, error)

var ErrPoolClosed = fmt.Errorf("Pool is closed")

type Pool struct {
	// protects the following fields
	mu          sync.Mutex
	numOpen     int // 当前打开的连接数
	MaxOpen     int // 允许打开的最大连接数
	MaxIdle     int // 允许存在的最大空闲连接数
	freeConn    []*poolConn
	connReq     map[int64]chan *poolConn
	nextReqKey  int64
	connFactory ConnFactory
	closed      bool
}

func NewPool() *Pool {
	return &Pool{}
}

func (p *Pool) nextReqKeyLocked() int64 {
	p.nextReqKey++
	return p.nextReqKey
}

func (p *Pool) Get(ctx context.Context) (*poolConn, error) {
	p.mu.Lock()
	if p.closed {
		p.mu.Unlock()
		return nil, ErrPoolClosed
	}

	// 若还有空闲连接，直接从freeConn拿来返回即可
	if len(p.freeConn) > 0 {
		p.numOpen++
		conn := p.freeConn[0]
		// p.freeConn = p.freeConn[1:]
		// 以下两行的写法跟上面那行的写法左右一样，都是移除slice中的第一个元素，只不过下面的写法会复用空间
		copy(p.freeConn, p.freeConn[1:])
		p.freeConn = p.freeConn[:len(p.freeConn)-1]

		p.mu.Unlock()
		return conn, nil
	}
	// 若没有空闲连接了而且已经打开了maxOpen的连接，则发送 connReq 等待连接回收
	if p.MaxOpen > 0 && p.numOpen >= p.MaxOpen {
		reqChan := make(chan *poolConn, 1)
		nextKey := p.nextReqKeyLocked()
		p.connReq[nextKey] = reqChan
		p.mu.Unlock()

		select {
		case conn, ok := <-reqChan:
			if !ok {
				return nil, ErrPoolClosed
			}
			return conn, nil
		case <-ctx.Done():
			p.mu.Lock()
			delete(p.connReq, nextKey)
			p.mu.Unlock()

			select {
			default:
			case conn, ok := <-reqChan:
				if ok && conn != nil {
					p.Put(ctx, conn)
				}
			}
			return nil, ctx.Err()
		}
	}
	// 若还没有超过 MaxOpen 则直接新建一个连接即可
	p.numOpen++
	p.mu.Unlock()
	ci, err := p.connFactory()
	if err != nil {
		p.mu.Lock()
		p.numOpen--
		p.mu.Unlock()
		return nil, err
	}
	pc := &poolConn{
		pool:      p,
		ci:        ci,
		createdAt: nowFunc(),
	}
	return pc, nil
}

var nowFunc = time.Now

// Put a conn back into the pool
func (p *Pool) Put(ctx context.Context, conn *poolConn) {
	p.mu.Lock()
	defer p.mu.Unlock()

	// 若 pool 已关闭，则直接将 conn 关闭即可
	if p.closed {
		conn.Close()
		p.numOpen--
		return
	}

	if p.MaxOpen > 0 && p.numOpen > p.MaxOpen {
		conn.Close()
		p.numOpen--
		return
	}

	// try to satisfy a connReq first
	if len(p.connReq) > 0 {
		var reqKey int64
		var req chan *poolConn
		for reqKey, req = range p.connReq {
			break
		}
		delete(p.connReq, reqKey)
		req <- conn
		return
	}
	// freeConn is aka idle conn
	if len(p.freeConn) >= p.MaxIdle {
		// close the conn
		conn.Close()
		p.numOpen--
		// todo: start a new goroutine to clean idle conn asyncly
		return
	} else {
		p.freeConn = append(p.freeConn, conn)
		return
	}
}

// Close close the pool
func (p *Pool) Close() error {
	p.mu.Lock()
	// make Close idempotent
	if p.closed {
		p.mu.Unlock()
		return nil
	}
	// 关闭空闲连接
	var err error
	fns := make([]func() error, 0, len(p.freeConn))
	for _, conn := range p.freeConn {
		fns = append(fns, conn.Close)
	}
	p.freeConn = nil
	p.closed = true
	// 关闭 connReq
	for _, req := range p.connReq {
		close(req)
	}
	// 先释放锁，再慢慢关闭空闲连接
	p.mu.Unlock()
	for _, fn := range fns {
		err1 := fn()
		if err1 != nil {
			err = err1
		}
	}
	return err
}
