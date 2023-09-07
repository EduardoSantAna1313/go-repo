package main

import "fmt"

func fib(max int) (int) {
	if (max == 0) {
		return 0
	} else if (max == 1) {
		return 1;
	} else if (max == 2) {
		return 1;
	}
	return fib(max-1) + fib(max -2)
}

func main() {
	fmt.Println("Result: ", fib(0)) // 0
	fmt.Println("Result: ", fib(1)) // 1
	fmt.Println("Result: ", fib(2)) // 1
	fmt.Println("Result: ", fib(3)) // 2
	fmt.Println("Result: ", fib(4)) // 3
	fmt.Println("Result: ", fib(5)) // 5
	fmt.Println("Result: ", fib(6)) // 8
	fmt.Println("Result: ", fib(7)) // 13
	fmt.Println("Result: ", fib(8)) // 21
	fmt.Println("Result: ", fib(9)) // 34
	fmt.Println("Result: ", fib(10)) // 55
	fmt.Println("Result: ", fib(20)) // 6765
	fmt.Println("Result: ", fib(30)) // 832040
	fmt.Println("Result: ", fib(40)) // 102334155
	fmt.Println("Result: ", fib(45)) // 1134903170
	fmt.Println("Result: ", fib(50)) // 12586269025
}