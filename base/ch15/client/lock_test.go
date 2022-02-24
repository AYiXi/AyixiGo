package client

import (
	"sync"
	"testing"
	"time"
)

func Test(t *testing.T) {
	var mut sync.Mutex
	c := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()

			mut.Lock()
			c++
		}()
	}

	time.Sleep(1 * time.Second)
	t.Logf("c = %d", c)
}

func Test2(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	c := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()

			mut.Lock()
			c++
			wg.Done()
		}()
	}

	wg.Wait()
	t.Logf("c = %d", c)
}
