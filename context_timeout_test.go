package context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContextWithTimeout(t *testing.T) {
	fmt.Println("Total goroutine", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	destination := CreateCounterWithContext(ctx)
	fmt.Println("Total goroutine", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("counter", n)
	}
	time.Sleep(2 * time.Second)

	fmt.Println("Total goroutine", runtime.NumGoroutine())
}
