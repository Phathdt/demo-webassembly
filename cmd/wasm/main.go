package main

import (
	"fmt"
	"syscall/js"
	"time"
)

var memo = make(map[int]int)

func fibo(n int) int {
	// Check if the result is already memoized
	if val, ok := memo[n]; ok {
		return val
	}

	// Base cases
	if n <= 1 {
		memo[n] = n
		return n
	}

	// Compute the Fibonacci number recursively
	memo[n] = fibo(n-1) + fibo(n-2)

	return memo[n]
}

func jsonWrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}

		n := args[0].Int()
		start := time.Now()

		value := fibo(n)

		end := time.Now()

		fmt.Println("value", value)
		fmt.Println("take", end.Sub(start))

		return value
	})

	return jsonFunc
}

func main() {
	fmt.Println("Go Web Assembly")
	js.Global().Set("fibo", jsonWrapper())
	<-make(chan bool)
}
