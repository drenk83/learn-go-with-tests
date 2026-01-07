package sum

func Sum(nums [5]int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}
	return sum
}

func SumSlice(nums []int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}
	return sum
}

func SumAll(nums ...[]int) []int {
	out := make([]int, len(nums))

	for i := range nums {
		out[i] = SumSlice(nums[i])
	}

	return out
}

func SumAllTails(nums ...[]int) []int {
	out := make([]int, len(nums))

	for i, numbers := range nums {
		if len(numbers) > 1 {
			out[i] = SumSlice(numbers[1:])
		} else {
			out[i] = 0
		}
	}

	return out
}