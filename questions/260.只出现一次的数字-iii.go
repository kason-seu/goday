package questions
/*
 * @lc app=leetcode.cn id=260 lang=golang
 *
 * [260] 只出现一次的数字 III
 */

// @lc code=start
func singleNumber(nums []int) []int {


	m := make(map[int]int)
	for _, val := range nums {

		m[val] = m[val] + 1
		
	}

	re := make([]int, 2)

	i := 0
	for k, v := range m {
		if v == 1 {
			re[i] = k
			i++
		}
	}

	return re
}
// @lc code=end

