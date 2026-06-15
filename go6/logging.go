package main

import (
	"context"
	"fmt"
	"time"
)

type LoggingService struct {
	next Service
}

func NewLoggingService(next Service) Service {
	return &LoggingService{next: next}
}

func (s *LoggingService) GetCatFact(ctx context.Context) (fact *CatFact, err error) {
	// defer function to log the time taken for the request, will be executed after the function returns, and will have access to the named return values fact and err
	defer func(start time.Time) {
		fmt.Printf("fact=%v, err=%v, took=%v\n", fact.Fact, err, time.Since(start))
	}(time.Now())

	return s.next.GetCatFact(ctx)
}
