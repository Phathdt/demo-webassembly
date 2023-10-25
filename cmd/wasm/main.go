package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"syscall/js"
	"time"
)

func fibo(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	return fibo(n-1) + fibo(n-2)
}

func jsonWrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}

		inputJSON := args[0].String()
		start := time.Now()

		log.WithFields(log.Fields{"start": start}).Info("A walrus appears")

		value := fibo(40)

		end := time.Now()

		log.WithFields(log.Fields{"end": end}).Info("A walrus appears")

		log.WithFields(log.Fields{"take": end.Sub(start)}).Info("A walrus appears")

		fmt.Println("value", value)

		return inputJSON
	})
	return jsonFunc
}

func main() {
	fmt.Println("Go Web Assembly")
	js.Global().Set("formatJSON", jsonWrapper())
	<-make(chan bool)
}
