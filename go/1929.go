package _go

func getConcatenation(nums []int) []int {
	ans := make([]int, 0, len(nums)*2)
	ans = append(ans, nums...)
	ans = append(ans, nums...)
	return ans
}