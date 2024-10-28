package main

import "time"

func main() {
	serverDone := make(chan struct{})
	var ticker = time.NewTicker(1 * time.Millisecond)
	defer ticker.Stop()
	var counter = 0
	for {
		select {
		case <-serverDone:
			return
		case <-ticker.C:
			counter += 1
		}
	}
}
