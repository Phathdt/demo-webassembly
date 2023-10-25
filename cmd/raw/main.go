package main

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

func fibo(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	return fibo(n-1) + fibo(n-2)
}

func main() {
	start := time.Now()

	log.WithFields(log.Fields{"start": start}).Info("A walrus appears")

	value := fibo(40)

	end := time.Now()

	log.WithFields(log.Fields{"end": end}).Info("A walrus appears")

	log.WithFields(log.Fields{"take": end.Sub(start)}).Info("A walrus appears")

	fmt.Println("value", value)
}
