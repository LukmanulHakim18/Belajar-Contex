package context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func CreateCounter() chan int {
	destination := make(chan int)
	go func() {

		counter := 1
		for {
			destination <- counter
			counter++
		}
	}()
	return destination
}

func TestGoroutineLeakes(t *testing.T) {
	fmt.Println("Total goroutine", runtime.NumGoroutine())
	destination := CreateCounter()
	for n := range destination {
		fmt.Println("counter", n)
		if n == 10 {
			break
		}
	}

	fmt.Println("Total goroutine", runtime.NumGoroutine())
}

// context dengan signal cancel
func CreateCounterWithContext(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second)
			}

		}
	}()
	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total goroutine", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)
	destination := CreateCounterWithContext(ctx)
	fmt.Println("Total goroutine", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("counter", n)
		if n == 10 {
			break
		}
	}
	cancel()
	time.Sleep(2 * time.Second)

	fmt.Println("Total goroutine", runtime.NumGoroutine())
}
