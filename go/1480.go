package _go

func runningSum(nums []int) []int {
	var result []int
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			result = append(result, nums[i])
		} else {
			result = append(result, result[i-1]+nums[i])
		}
	}
	return result
}