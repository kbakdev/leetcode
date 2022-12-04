package _go

func pivotIndex(nums []int) int {
	var sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	var left int
	for i := 0; i < len(nums); i++ {
		if left == sum-nums[i]-left {
			return i
		}
		left += nums[i]
	}
	return -1
}
