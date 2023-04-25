package main

import (
	"log"
	"time"
)

func trace(msg string) func() {
	start := time.Now()
	log.Printf("вход в %s", msg)
	return func() {
		log.Printf("выход из %s (%s)", msg, time.Since(start))
	}
}

func bigSlowOperation() {
	defer trace("bigSlowOperation")()

	time.Sleep(2 * time.Second)
}

func main() {
	bigSlowOperation()
}
