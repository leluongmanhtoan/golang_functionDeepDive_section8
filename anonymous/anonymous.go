package anonymous

import "fmt"

//type anotherFN func(int, []string, map[string][]int) ([]int, string)	Trường hợp khi phải truyền vào quá dài khiến code trở nên khó maintain

func Anonymous() {
	numbers := []int{1, 2, 3}
	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})
	double := createTransformer(2)
	fmt.Println(transformed)
	doubled := transformNumbers(&numbers, double)
	fmt.Println(doubled)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))

	}
	return dNumbers
}

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
