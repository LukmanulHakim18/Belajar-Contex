package context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContextWithDeadline(t *testing.T) {
	fmt.Println("Total goroutine", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(10*time.Second))
	defer cancel()

	destination := CreateCounterWithContext(ctx)
	fmt.Println("Total goroutine", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("counter", n)
	}

	fmt.Println("Total goroutine", runtime.NumGoroutine())
}
