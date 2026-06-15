package main

import (
	"log"
)

func main() {
	svc := NewCatFactService("https://catfact.ninja/fact")
	svc = NewLoggingService(svc)

	apiServer := NewApiServer(svc)
	if err := apiServer.Start(":8080"); err != nil {
		log.Fatal(err)
	}

	// fact, err := svc.GetCatFact(context.TODO())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("Cat Fact: %s\n", fact.Fact)
}
