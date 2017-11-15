package main

import (
	"fmt"
	"os"
	"time"

	"./semaphore"
)

func main() {
	semaphoreWithTimeouts()
	semaphoreWithoutTimeouts()
}

func semaphoreWithTimeouts() {
	tickets, timeout := 1, 3*time.Second
	s := semaphore.New(tickets, timeout)

	if err := s.Acquire(); err != nil {
		panic(err)
	}

	// Do important work
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}

	if err := s.Release(); err != nil {
		panic(err)
	}
}

func semaphoreWithoutTimeouts() {
	tickets, timeout := 0, 0*time.Second
	s := semaphore.New(tickets, timeout)

	if err := s.Acquire(); err != nil {
		if err != semaphore.ErrNoTickets {
			panic(err)
		}

		// No tickets left, can't work :(
		os.Exit(1)
	}
}
