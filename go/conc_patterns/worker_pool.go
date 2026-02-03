package concpatterns

import (
	"context"
	"fmt"
	"sync"
)

const workers = 8

func WorkersPool() {
	// inCh := make(chan int, workers)
	inCh := Generator(100) // out generator as input channel
	outCh := make(chan int, 100)

	// for preventing goroutine leak
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	wg := sync.WaitGroup{}
	wg.Add(workers)

	for range workers { // launch N workers

		// TODO: add dynamic workerpool (optional)

		// Worker pool -- is fan-out + restrict count of workers exist
		// classic fan-out doesn't restrict number of workers
		go func(ctx context.Context, inCh <-chan int, outCh chan<- int) {
			// each worker receive value and square it again
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case e, ok := <-inCh:
					if !ok {
						return
					}
					fmt.Printf("received '%d'\n", e)
					outCh <- e * e
				}
			}
		}(ctx, inCh, outCh)
	}

	// wait until receive all results
	wg.Wait()
	close(outCh)
	// if we want to receive values imeadently -- hide two previous Row into goroutine
	/*
		go func(){
			wg.Wait()
			close(outCh)
		}()
	*/

	for out := range outCh {
		fmt.Println(out)
	}

}
