package main

import "fmt"

type Person struct {
	age int
}

// 通常我们使用指针类型作为方法的接收者的理由:
// 1.使用指针类型能够修改调用者的值
// 2.使用指针类型可以避免在每次调用方法时复制该值，在值的类型为大型结构体时，这样做会更加高效，
func main() {
	p := Person{age: 10}

	p.IncrAge()
	fmt.Printf("IncrAge value is %d \n", p.GetAge())

	p.PointerIncrAge()
	fmt.Printf("PointerIncrAge value is %d \n", p.GetAge())

	fmt.Printf("p.age value is %d \n", p.age)

	fmt.Println()
	p2 := &Person{age: 20}

	p2.IncrAge()
	fmt.Printf("IncrAge value is %d \n", p2.GetAge())

	p2.PointerIncrAge()
	fmt.Printf("PointerIncrAge value is %d \n", p2.GetAge())

	fmt.Printf("p.age value is %d \n", p2.age)

}

// IncrAge 如果实现了接收者是指针类型的方法，会隐含地也实现了接收者是值类型的 IncrAge 方法。
// 会修改age的值
func (p Person) IncrAge() {
	p.age++
}

// PointerIncrAge 如果实现了接收者是指针类型的方法，会隐含地也实现了接收者是值类型的 PointerIncrAge 方法
// 会修改age的值
func (p *Person) PointerIncrAge() {
	p.age++
}

// GetAge 如果实现了接收者是值类型的方法，会隐含地也实现了接收者是指针类型的 GetAge 方法。
func (p Person) GetAge() int {
	return p.age
}
