package main

import (
	"context"
	"net"
)

// 参考自：https://www.youtube.com/watch?v=5zXAHh5tJqQ

type token struct{}

type PoolSema struct {
	sem  chan token
	idle chan net.Conn
}

func NewPoolSema(limit int) *PoolSema {
	sem := make(chan token, limit)
	idle := make(chan net.Conn, limit)

	return &PoolSema{sem, idle}
}

func (p *PoolSema) Release(c net.Conn) {
	p.idle <- c
}

func (p *PoolSema) Hijack(c net.Conn) {
	<-p.sem
}

func dial() (net.Conn, error) {
	return nil, nil
}

func (p *PoolSema) Accquire(ctx context.Context) (net.Conn, error) {
	select {
	case conn := <-p.idle:
		return conn, nil
	case p.sem <- token{}:
		conn, err := dial()
		if err != nil {
			<-p.sem
		}
		return conn, err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
