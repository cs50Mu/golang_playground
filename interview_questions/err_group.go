package main

import "sync"

type ErrGroup struct {
	wg   sync.WaitGroup
	once sync.Once
	err  error
}

func (eg *ErrGroup) Do(f func() error) {
	eg.wg.Add(1)
	go func() {
		defer eg.wg.Done()
		if err := f(); err != nil {
			eg.once.Do(func() {
				eg.err = err
			})
		}
	}()
}

func (eg *ErrGroup) Wait() error {
	eg.wg.Wait()
	return eg.err
}
