package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    *sync.Mutex
	wg    *sync.WaitGroup
	count int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.count += 10
	c.wg.Done()
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.count
}

func main() {

	c := Counter{
		mu:    &sync.Mutex{},
		count: 0,
		wg:    &sync.WaitGroup{},
	}

	fmt.Printf("启动 100 个协程,每个协程累加 10 次...\n")

	for i := 0; i < 100; i++ {
		c.wg.Add(1)
		go c.Increment()
	}

	c.wg.Wait()
	fmt.Printf("最终计数: %v\n", c.count)

}
