package functionarevalues

import "fmt"

type transformFn func(int) int

//type anotherFN func(int, []string, map[string][]int) ([]int, string)	Trường hợp khi phải truyền vào quá dài khiến code trở nên khó maintain

func main() {
	numbers := []int{1, 2, 3, 4}
	tripled := transformNumbers(&numbers, triple)
	fmt.Println(tripled)
	double := transformNumbers(&numbers, double)
	fmt.Println(double)
	moreNumber := []int{5, 1, 2}
	transformFn1 := getTransformerFunction(&numbers)
	transformFn2 := getTransformerFunction(&moreNumber)
	transformedNumbers := transformNumbers(&numbers, transformFn1)

	moreTransformedNumbers := transformNumbers(&moreNumber, transformFn2)
	fmt.Println(transformedNumbers)
	fmt.Println(moreTransformedNumbers)
}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))

	}
	return dNumbers
}

func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}

}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
