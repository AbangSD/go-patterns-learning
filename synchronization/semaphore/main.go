package main

import (
	"fmt"
	"os"
	"time"

	"github.com/AbangSD/go-patterns-learning/synchronization/semaphore/semaphore"
)

func main() {
	fmt.Println("test one")
	semaphoreWithTimeouts1()

	fmt.Println("test two")
	semaphoreWithTimeouts2()

	fmt.Println("test three")
	semaphoreWithoutTimeouts()
}

func semaphoreWithTimeouts1() {
	tickets, timeout := 3, 3*time.Second
	s := semaphore.New(tickets, timeout)

	go func() {
		for i := 0; i < 5; i++ {
			if err := s.Acquire(); err != nil {
				fmt.Println("error", err, time.Now())
			} else {
				fmt.Print("no error", time.Now())
				// Do important work here
				fmt.Println(" something important")
			}
		}
	}()

	time.Sleep(4 * time.Second)

	if err := s.Release(); err != nil {
		panic(err)
	}

	time.Sleep(5 * time.Second)
}

func semaphoreWithTimeouts2() {
	tickets, timeout := 3, 3*time.Second
	s := semaphore.New(tickets, timeout)

	go func() {
		for i := 0; i < 5; i++ {
			if err := s.Acquire(); err != nil {
				fmt.Println("error", err, time.Now())
			} else {
				fmt.Print("no error", time.Now())
				// Do important work here
				fmt.Println(" something important")
			}
		}
	}()

	time.Sleep(2 * time.Second)

	if err := s.Release(); err != nil {
		panic(err)
	}

	time.Sleep(5 * time.Second)
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
