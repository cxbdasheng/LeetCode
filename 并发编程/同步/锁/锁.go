package main

import "sync"

func main() {
	//var mutex sync.Mutex
	//defer mutex.Unlock()
	//mutex.Unlock()

	var rwm sync.RWMutex
	rwm.RUnlock()
}
