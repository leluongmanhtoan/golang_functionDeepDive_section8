package recursion

func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}

func fibonacci(number int) int {
	if number <= 1 {
		return number
	}
	return fibonacci(number-1) + fibonacci(number-2)
}
