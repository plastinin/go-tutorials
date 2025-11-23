package main

import (
	"context"
	"fmt"
	"time"
)

// родительский контекст
func foo(ctx context.Context, n int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("foo done.", n)
			return
		default:
			fmt.Println("foo", n)
		}
		time.Sleep(100 * time.Millisecond)
	}
}

// дочерний контекст
func boo(ctx context.Context, n int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("boo done.", n)
			return
		default:
			fmt.Println("boo", n)
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {

	parentContent, parentCancel := context.WithCancel(context.Background())
	childContent, childCancel := context.WithCancel(parentContent)

	go foo(parentContent, 1)
	go foo(parentContent, 2)
	go foo(parentContent, 3)

	go boo(childContent, 1)
	go boo(childContent, 2)
	go boo(childContent, 3)
	
	time.Sleep(1 * time.Second)
	childCancel()
	
	time.Sleep(1 * time.Second)
	parentCancel()
	time.Sleep(3 * time.Second)
	fmt.Println("main done.")

}
