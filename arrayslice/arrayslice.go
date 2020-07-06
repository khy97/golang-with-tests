package arrayslice

// Sum -> sums the elements of an array
func Sum(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}

	return sum
}

// SumAll -> Returns the sum of all slices in a slice
func SumAll(numbersToSum ...[]int) (sums []int) {
	sums = make([]int, len(numbersToSum))

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return
}

// SumAllTails -> Returns the sum of all slices in a slice excluding the first character
func SumAllTails(numbersToSum ...[]int) (sums []int) {
	sums = make([]int, len(numbersToSum))

	for i, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums[i] = 0
			continue
		}
		sums[i] = Sum(numbers[1:])
	}
	return
}
