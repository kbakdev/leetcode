package _go

func getConcatenation(nums []int) []int {
	var result []int
	for i := 0; i < len(nums); i++ {
		result = append(result, nums[i])
	}
	for i := 0; i < len(nums); i++ {
		result = append(result, nums[i])
	}
	return result
}