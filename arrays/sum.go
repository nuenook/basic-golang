package arrays

import "fmt"

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAllTails(numbersToSum ...[]int) []int {
	// lengthOfNumbers := len(numbersToSum)
	// sums := make([]int, lengthOfNumbers)

	// for i, numbers := range numbersToSum {
	// 	fmt.Println(numbers)
	// 	sums[i] = Sum(numbers)
	// }

	// way 2...
	// var sums []int
	// for _, numbers := range numbersToSum {
	// 	sums = append(sums, Sum(numbers))

	// }

	// Sum Tails
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	fmt.Println(sums)
	return sums
}
