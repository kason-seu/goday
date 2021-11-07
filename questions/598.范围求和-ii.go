package questions

/*
 * @lc app=leetcode.cn id=598 lang=golang
 *
 * [598] 范围求和 II
 */

// @lc code=start
func maxCount(m int, n int, ops [][]int) int {

	// 求交集
	rowMin, colMin := m, n
	for _, row := range ops {

		
		if rowMin > row[0] {
			rowMin = row[0]
		}

		
		if colMin > row[1] {
			colMin = row[1]
		}

	}

	return rowMin * colMin
}

// @lc code=end
