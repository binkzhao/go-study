package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//useWithTimeout()
	useWithValue()
}

// WithTimeout 返回 WithDeadline(parent, time.Now().Add(timeout))。
func useWithTimeout() {
	// Pass a context with a timeout to tell a blocking function that it
	// should abandon its work after the timeout elapses.
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println("Msg: ", ctx.Err()) // prints "context deadline exceeded"
	}
}

func useWithDeadline() {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// Even though ctx will be expired, it is good practice to call its
	// cancelation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println("Msg: ", ctx.Err())
	}
}

func useWithValue() {
	type contextKey string
	f := func(ctx context.Context, k contextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("Found Value: ", v)
			return
		}
		fmt.Println("Key Not Found: ", k)
	}

	k := contextKey("language")
	ctx := context.WithValue(context.Background(), k, "Go")
	f(ctx, k)
	f(ctx, contextKey("PHP"))
}
