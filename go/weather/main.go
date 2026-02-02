package main

import (
	"context"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

/*
study example how to concurent get wheater and return values in highload, a lot of read requests
*/

// don't edit START
var providerTime = 1 * time.Minute

func WeatherProvider(_ string) int {
	// _ -- any city
	time.Sleep(providerTime)

	return rand.IntN(100)
}

// don't edit END

var cities = []string{"Moscow", "St. Perersburg"} // and etc... this cities also can be got from some source
var wCache *WeatherCache

type WeatherCache struct {
	temperature map[string]int
	mu          sync.RWMutex
}

func NewWeatherCache(ctx context.Context, pollPeriod time.Duration, cities []string) *WeatherCache {
	ticker := time.NewTicker(pollPeriod)
	wCachePtr := &WeatherCache{
		temperature: make(map[string]int),
	}
	// init cache
	for _, city := range cities {
		wCachePtr.temperature[city] = 0
	}
	log.Println("start warm-up cache")
	wCachePtr.UpdateTemperature(ctx) // WARM UP data before starting server
	log.Println("warm-up cache done")

	// background update wheater values in some period
	go func() {
		defer ticker.Stop() // w/o leakign resources
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				wCachePtr.UpdateTemperature(ctx)
			}
		}
	}()

	return wCachePtr
}

func (w *WeatherCache) UpdateTemperature(ctx context.Context) {
	// background update for cities in map
	wg := sync.WaitGroup{}
	w.mu.RLock() // for excluding concurent map iteration panic
	for city := range w.temperature {
		wg.Add(1)
		go func(city string) {
			defer wg.Done()

			// don't start any new requests in case of cancel
			select {
			case <-ctx.Done():
				return
			default:
			}

			res := WeatherProvider(city)
			w.mu.Lock()
			w.temperature[city] = res
			w.mu.Unlock()
		}(city)
	}
	w.mu.RUnlock()
	wg.Wait()
}

func (w *WeatherCache) GetTemperature(city string) (int, error) {
	w.mu.RLock()
	defer w.mu.RUnlock()
	val, ok := w.temperature[city]
	if !ok {
		return 0, fmt.Errorf("no data for '%s' city", city)
	}

	return val, nil

}

func WeatherHandler(wCache *WeatherCache) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		qVals := r.URL.Query()
		city := qVals.Get("city")
		log.Printf("request received: city=%s\n", city)
		val, err := wCache.GetTemperature(city)
		if err != nil {
			log.Printf("request failed: unknown city=%s\n", city)
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}
		log.Printf("request ok: city=%s,val=%d\n", city, val)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(strconv.Itoa(val)))
	}

}

func main() {
	// catch signal for graceful shutdown
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	wCache = NewWeatherCache(ctx, providerTime, cities) // init global cache

	mux := http.NewServeMux()
	mux.HandleFunc("/weather", WeatherHandler(wCache))
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		log.Println("server started on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen error: %v", err)
		}
	}()

	<-ctx.Done() // wait for signal
	log.Println("shutting down...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_ = server.Shutdown(shutdownCtx)
	log.Println("server stopped")
}
