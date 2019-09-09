package main

import "fmt"

func main() {
	fmt.Println(factorial(20))

	fmt.Println(fib(7))
}

func factorial(n int) int {
	if n == 1 || n == 0 {
		return 1
	}

	//f(n) = f(n - 1)*n
	return factorial(n-1) * n
}

func fib(n int) int {
	if n == 1 || n == 2{
		return 1
	}

	return fib(n - 1) + fib(n-2)
}

