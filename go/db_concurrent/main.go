package dbconcurrent

import (
	"context"
	"errors"
	"sync"
	"time"
)

var ErrNotFound = errors.New("not found")
var ErrNetworkIssue = errors.New("network issue")
var ErrTimeout = errors.New("timeout")

const retry = 3 // for SQL requests
const timeout = 5 * time.Second

type Response struct {
	Value string
	Err   error
}

type DatabaseHost interface {
	DoQuery(query string) (string, error)
}

func DistributedQuery(query string, replicas []DatabaseHost) (string, error) {
	resCh := make(chan Response, len(replicas)) // make buffered for non-blocking write for all subgoroutines

	ctx, cancelFunc := context.WithTimeout(context.Background(), timeout)
	defer cancelFunc()

	wg := sync.WaitGroup{}
	wg.Add(len(replicas))
	for _, rep := range replicas {
		// make parallel SQL requests to all replic as
		go func(ctx context.Context, replica DatabaseHost) {
			defer wg.Done()
			for range retry {
				if ctx.Err() != nil {
					// request cancelled
					// compact view if '<- ctx.Done()'
					// 1. timeout deadline exceed
					// 2. not found or resut recevied
					break
				}

				res, err := replica.DoQuery("SELECT * FROM test;")
				if err == nil {
					if res == "" {
						// empty response == not found
						resCh <- Response{
							"",
							ErrNotFound,
						}
					} else {
						// smth received
						resCh <- Response{
							res,
							err,
						}
					}
					return
				}
				// wait before retry
				<-time.After(time.Second) // NOTE: по хорошему тут должен быть backoff, но упростим

			}
			// no answer after retry
			// no response in channel

		}(ctx, rep)
	}

	go func() {
		// close channel after receiveng all answers
		wg.Wait()
		close(resCh)
	}()

	for {
		select {
		case res, ok := <-resCh:
			if !ok {
				return "", ErrNetworkIssue
			}
			if res.Err != nil {
				cancelFunc()
				return "", ErrNotFound
			}
			cancelFunc()
			return res.Value, nil
		case <-ctx.Done():
			// timeout on all requests
			return "", ErrTimeout
		}

	}
}
