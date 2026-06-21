package main

import (
	"flag"
	"fmt"
	"log"

	"sc/go13/api"
	"sc/go13/storage"
)

func main() {
	listenAddr := flag.String("listen", ":3000", "http listen address")
	flag.Parse()

	store := storage.NewMemoryStorage() // Use MemoryStorage for simplicity

	server := api.NewServer(*listenAddr, store)
	fmt.Println("Server is running on", *listenAddr)
	log.Fatal(server.Start())
}
