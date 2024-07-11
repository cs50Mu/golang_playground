package main

import (
	"fmt"
	"log"
	"testing"
)

func TestErrGroup(t *testing.T) {
	var eg ErrGroup
	for i := 0; i < 100; i++ {
		i := i
		eg.Do(func() error {
			fmt.Printf("doing %v\n", i)
			if i > 90 {
				return fmt.Errorf("error at %v", i)
			} else {
				return nil
			}
		})
	}

	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}
}
