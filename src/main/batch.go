package main

// Background task thread.
// SingleTaskQueue for heavy works.

import (
	"sync"
	"time"
)

var batchChannel chan<- func()

func initializeBatch() {
	var rc chan func()
	rc, batchChannel = createChannels()
	go runWorker(rc)
}

func createChannels() (chan func(), chan<- func()) {
	receiver := make(chan func())
	return receiver, receiver
}

// when AddToBachQueue is called and added function,
// this go routine executes it one by one.
func runWorker(receiver chan func()) {
	var wg sync.WaitGroup
	defer close(receiver)
	wg.Add(1)
	go func() {
		for {
			select {
			case f := <-receiver:
				f()
			default:
				time.Sleep(100 * time.Millisecond)
			}
		}
		wg.Done()
	}()
	wg.Wait()
}

func AddToBatchQueue(action func()) {
	batchChannel <- action
}

func terminateBatch() {
	close(batchChannel)
}
