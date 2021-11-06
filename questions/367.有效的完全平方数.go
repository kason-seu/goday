package questions
/*
 * @lc app=leetcode.cn id=367 lang=golang
 *
 * [367] 有效的完全平方数
 */

// @lc code=start
func isPerfectSquare(num int) bool {
	i,j:=0,num
	
	for; i <= j; {
		mid := (i + j) / 2
		if mid * mid == num {
			return true
		} else if mid * mid < num {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return false
}
// @lc code=end

