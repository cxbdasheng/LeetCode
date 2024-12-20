package main

import (
	"fmt"
	"sync"
	"time"
)

type Queue struct {
	items []int
	mu    sync.Mutex
	cond  *sync.Cond
}

func NewQueue() *Queue {
	q := &Queue{}
	q.cond = sync.NewCond(&q.mu)
	return q
}

// 生产者
func (q *Queue) Produce(item int) {
	q.mu.Lock()
	q.items = append(q.items, item)
	fmt.Printf("生产: %d\n", item)
	q.cond.Signal() // 通知消费者
	q.mu.Unlock()
}

// 消费者
func (q *Queue) Consume() {
	q.mu.Lock()
	for len(q.items) == 0 {
		q.cond.Wait() // 等待条件
	}
	item := q.items[0]
	q.items = q.items[1:]
	fmt.Printf("消费: %d\n", item)
	q.mu.Unlock()
}

func main() {

	queue := NewQueue()
	var wg sync.WaitGroup

	// 启动消费者
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			queue.Consume()
			time.Sleep(time.Millisecond * 500)
		}
	}()

	// 启动生产者
	for i := 0; i < 5; i++ {
		queue.Produce(i)
		time.Sleep(time.Millisecond * 300)
	}

	wg.Wait() // 等待消费者完成
}
