package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var i32 int32
	i32 = 32
	atomic.AddInt32(&i32, 1)
	fmt.Println(i32)
	atomic.AddInt32(&i32, -2)
	fmt.Println(i32)

	var ui32 uint32 = 32
	delta := int32(-3)
	atomic.AddUint32(&ui32, uint32(delta))
	fmt.Println(ui32)

	atomic.AddUint32(&ui32, ^uint32(-(-3)-1))
	fmt.Println(ui32)

	var num int32 = 32
	atomic.CompareAndSwapInt32(&num, 32, 1)
	fmt.Println(num)

	var num2 int32 = 32
	atomic.StoreInt32(&num2, 31)

	//var num2 int32 = 20
	//for {
	//	v := atomic.LoadInt32(&num2)
	//	if atomic.CompareAndSwapInt32(&num2, v, 0) {
	//		fmt.Println("The second number has gone to zero.")
	//		break
	//	}
	//	time.Sleep(time.Millisecond * 500)
	//}
}
