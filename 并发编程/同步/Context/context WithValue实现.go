package main

import (
	"context"
	"fmt"
)

func step1(ctx context.Context) context.Context {
	child := context.WithValue(ctx, "name", "陈大剩")
	return child
}

func step2(ctx context.Context) context.Context {
	child := context.WithValue(ctx, "age", "18")
	return child
}
func step3(ctx context.Context) {
	fmt.Printf("name %s \n", ctx.Value("name"))
	fmt.Printf("age %s \n", ctx.Value("age"))
}

func main() {
	parent := context.Background()
	child1 := step1(parent)
	child2 := step2(child1)
	step3(child2)
}
