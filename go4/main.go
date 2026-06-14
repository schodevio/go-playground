package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	// ctx := context.Background() // ctx := context.TODO() // both are same, but TODO is used when you are not sure which one to use
	ctx := context.WithValue(context.Background(), "foo", "bar") // adding value to context
	userID := 123

	val, err := fetchUserData(ctx, userID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Fetched value:", val)
	fmt.Println("Time taken:", time.Since(start))
}

type Response struct {
	value int
	err   error
}

func fetchUserData(ctx context.Context, userID int) (int, error) {
	foo := ctx.Value("foo")
	fmt.Println("Value of foo in context:", foo)

	ctx, cancel := context.WithTimeout(ctx, 400*time.Millisecond) // setting timeout for the context, after which it will be cancelled automatically
	defer cancel()

	respch := make(chan Response)

	go func() {
		fmt.Println("Fetching data for user:", userID)

		val, err := fetchThirdPartyStuffSlowly()
		respch <- Response{
			value: val,
			err:   err,
		}
	}()

	for {
		select {
		case <-ctx.Done(): // this case will be executed when the context is cancelled, either due to timeout or manually
			return 0, fmt.Errorf("fetching data for user %d timed out", userID)
		case resp := <-respch: // this case will be executed when the response is received from the goroutine
			return resp.value, resp.err
		}
	}
}

func fetchThirdPartyStuffSlowly() (int, error) {
	time.Sleep(500 * time.Millisecond)
	return 666, nil
}
