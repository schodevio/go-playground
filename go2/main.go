package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	userName := fetchUser()

	respch := make(chan any, 2) // buffered channel to avoid blocking
	wg := &sync.WaitGroup{}     // wait group to wait for all goroutines to finish
	wg.Add(2)                   // we have two goroutines to wait for

	go fetchUserLikes(userName, respch, wg)
	go fetchUserMatch(userName, respch, wg)

	wg.Wait()     // wait for both goroutines to finish
	close(respch) // close the channel after all goroutines have finished sending data

	fmt.Println("------------------------------")

	for resp := range respch {
		fmt.Printf("Received response: %v\n", resp)
	}

	// fmt.Printf("User %s has %d likes and a match with %s\n", userName, likes, match)
	fmt.Printf("Execution time: %s\n", time.Since(start))
}

func fetchUser() string {
	fmt.Println("Fetching user...")

	time.Sleep(100 * time.Millisecond)
	return "John Doe"
}

func fetchUserLikes(userName string, respch chan any, wg *sync.WaitGroup) {
	defer wg.Done() // mark this goroutine as done when it finishes

	fmt.Printf("Fetching likes for user %s...\n", userName)

	time.Sleep(200 * time.Millisecond)
	respch <- 42 // sending the number of likes to the channel
}

func fetchUserMatch(userName string, respch chan any, wg *sync.WaitGroup) {
	defer wg.Done() // mark this goroutine as done when it finishes

	fmt.Printf("Fetching matches for user %s...\n", userName)

	time.Sleep(100 * time.Millisecond)
	respch <- "Alice" // sending the match name to the channel
}

//================================

// func main() {
// 	start := time.Now()
// 	userName := fetchUser()

// 	likes := fetchUserLikes(userName)
// 	match := fetchUserMatch(userName)

// 	fmt.Println("------------------------------")
// 	fmt.Printf("User %s has %d likes and a match with %s\n", userName, likes, match)
// 	fmt.Printf("Execution time: %s\n", time.Since(start))
// }

// func fetchUser() string {
// 	fmt.Println("Fetching user...")

// 	time.Sleep(100 * time.Millisecond)
// 	return "John Doe"
// }

// func fetchUserLikes(userName string) int {
// 	fmt.Printf("Fetching likes for user %s...\n", userName)

// 	time.Sleep(200 * time.Millisecond)
// 	return 42
// }

// func fetchUserMatch(userName string) string {
// 	fmt.Printf("Fetching matches for user %s...\n", userName)

// 	time.Sleep(100 * time.Millisecond)
// 	return "Alice"
// }
