package service

import (
	"log"
	"sync"
	"testing"
	_ "time"
)

func TestBatch(t *testing.T) {
	var wg sync.WaitGroup
	threadActions := make(chan func())
	defer close(threadActions)
	go runWorkerThread(threadActions)
	wg.Add(3)

	for i := 0; i < 3; i++ {
		go func() {
			threadActions <- func() {
				wg.Done()
			}
		}()
	}

	wg.Wait()
	t.Log(`Test Finished.`)
}

func runWorkerThread(queue chan func()) {
	for {
		job := <-queue
		job()
		log.Println(`work done`)
	}

}
