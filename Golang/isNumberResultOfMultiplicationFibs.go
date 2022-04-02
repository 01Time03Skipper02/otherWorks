package main

import (
	"fmt"
)

func ProductFib(prod uint64) [3]uint64 {
	fib := make([]uint64, 100)
	fib[0] = 1
	fib[1] = 1
	var res [3]uint64
	for i := 2; i < 100; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	for i := 1; prod > fib[i-1]*fib[i]; i++ {
		if prod == fib[i]*fib[i+1] {
			res[0] = fib[i]
			res[1] = fib[i+1]
			res[2] = 1
			fmt.Println(res[0], res[1])
			break
		}
		if prod < fib[i+1]*fib[i+2] {
			res[0] = fib[i+1]
			res[1] = fib[i+2]
			res[2] = 0
			break
		}
	}
	return res
}

func main() {
	fmt.Println(ProductFib(714))
}
