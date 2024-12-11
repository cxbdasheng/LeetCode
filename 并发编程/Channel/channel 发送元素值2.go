package main

import (
	"fmt"
	"time"
)

type Counter struct {
	count int
}

func main() {
	mapChan := make(chan map[string]*Counter, 1)
	syncChan := make(chan struct{}, 2)
	go func() {
		for {
			if elem, ok := <-mapChan; ok {
				counter := elem["count"]
				counter.count++
			} else {
				break
			}
		}
		fmt.Println("结束了接收")
		syncChan <- struct{}{}
	}()

	go func() {
		countMap := map[string]*Counter{
			"count": {},
		}
		for i := 0; i < 5; i++ {
			mapChan <- countMap
			time.Sleep(time.Millisecond) // 睡眠一微秒使得接收者有时间消费
			//fmt.Printf("这个 map 总数为：%v \n", countMap)
			fmt.Println(countMap["count"].String())
		}
		close(mapChan)
		syncChan <- struct{}{}
	}()
	<-syncChan
	<-syncChan
}

func (counter *Counter) String() string {
	return fmt.Sprintf("count:%d ", counter.count)
}
