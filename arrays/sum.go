package main

import "fmt"

// func Sum(numbers [5]int) int {
// 	sum := 0
// 	for i := 0; i < 5; i++ {
// 		sum += numbers[i]
// 	}
// 	return sum
// }

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// func SumAll(numbersToSum ...[]int) []int {
// 	lengthOfNumbers := len(numbersToSum)
// 	sums := make([]int, lengthOfNumbers)
// 	fmt.Println(sums)
// 	for i, numbers := range numbersToSum {
// 		sums[i] = Sum(numbers)
// 	}
// 	return sums
// }

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}

func main() {
	fmt.Println(SumAll([]int{2, 3, 4}, []int{1, 3, 5}))
}
