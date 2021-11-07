package questions
/*
 * @lc app=leetcode.cn id=268 lang=golang
 *
 * [268] 丢失的数字
 */

// @lc code=start
func missingNumber(nums []int) int {

	sum, len := 0,0
	for _, val := range nums {
		sum += val
		len += 1
	}

	total := ((len + 1) * len)/2

	return total - sum
}
// @lc code=end

