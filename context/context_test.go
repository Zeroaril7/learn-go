package context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContextValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	//Get Value
	fmt.Println(contextF.Value("c"))
	fmt.Println(contextA.Value("b"))
}

func TestContextCancel(t *testing.T) {

	parentCtx := context.Background()
	ctx, cancel := context.WithCancel(parentCtx)
	fmt.Println("Jumlah goroutine", runtime.NumGoroutine())

	destination := CreateCounter(ctx)
	fmt.Println("Jumlah goroutine", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println(n)
		if n == 10 {
			break
		}
	}
	cancel()
	time.Sleep(1 * time.Second)

	fmt.Println("Jumlah goroutine", runtime.NumGoroutine())
}

func TestContextTimeout(t *testing.T) {

	parentCtx := context.Background()
	ctx, cancel := context.WithTimeout(parentCtx, 5*time.Second)
	fmt.Println("Jumlah goroutine", runtime.NumGoroutine())
	defer cancel()

	destination := CreateCounterTimeOut(ctx)
	fmt.Println("Jumlah goroutine", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println(n)
	}

	fmt.Println("Jumlah goroutine", runtime.NumGoroutine())
}

func TestContextDeadline(t *testing.T) {

	parentCtx := context.Background()
	ctx, cancel := context.WithDeadline(parentCtx, time.Now().Add(7*time.Second))
	fmt.Println("Jumlah goroutine", runtime.NumGoroutine())
	defer cancel()

	destination := CreateCounterTimeOut(ctx)
	fmt.Println("Jumlah goroutine", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println(n)
	}

	fmt.Println("Jumlah goroutine", runtime.NumGoroutine())
}

func CreateCounter(ctx context.Context) chan int {
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
			}
		}
	}()

	return destination
}

func CreateCounterTimeOut(ctx context.Context) chan int {
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
