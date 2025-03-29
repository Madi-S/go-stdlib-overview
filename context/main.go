package main

import (
	"context"
	"fmt"
	"time"
)

type User struct {
	ID int
}

func main() {
	// Context with timeout
	// This context will be canceled after 2 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // Important: Always call `cancel` to release resources
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Task completed")
	case <-ctx.Done():
		fmt.Println("Timeout:", ctx.Err())
	}

	// Context with specific deadline
	// This context will be completed after 2 seconds
	deadline := time.Now().Add(5 * time.Second)
	ctx2, cancel2 := context.WithDeadline(context.Background(), deadline)
	defer cancel2()
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Task 2 completed")
	case <-ctx2.Done():
		fmt.Println("Deadline exceeded:", ctx2.Err())
	}

	// Context with cancel
	// This context will be canceled manually
	user := &User{ID: 42}
	ctx3, cancel3 := context.WithCancel(context.WithValue(context.TODO(), "user", user))
	go func() {
		ctx3.Value("user").(*User).ID = 100
		time.Sleep(2 * time.Second)
		cancel3()
	}()
	<-ctx3.Done() // Blocks until canceled
	fmt.Println("Canceled:", ctx3.Err(), ctx3.Value("user").(*User).ID)

	// Context with value (e.g., request-scoped data, not for global state)
	ctx4 := context.WithValue(context.Background(), "userID", 5)
	userID := ctx4.Value("userID")
	nonExisting := ctx4.Value("nonExistingKey")
	fmt.Println("User ID:", userID)
	fmt.Println("Non-existing key:", nonExisting)

	// Context in goroutines
	ctx5, cancel5 := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel5()

	go worker(ctx5) // goroutines respect `context.Done()`

	time.Sleep(4 * time.Second)
	fmt.Println("Main function exiting")
}

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker stopped:", ctx.Err())
			return
		default:
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
