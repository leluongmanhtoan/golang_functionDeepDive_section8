package main

import "fmt"

func main() {
	numbers := []int{1, 10, 10, 10}
	sum := sumup(1, numbers...)
	//option variadic
	sum2 := sumup(2, 10, 20, 30)
	fmt.Println(sum)
	fmt.Println(sum2)
}

func sumup(starting int, numbers ...int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}
