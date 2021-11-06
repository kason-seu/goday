package questions

import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=869 lang=golang
 *
 * [869] 重新排序得到 2 的幂
 */

// @lc code=start
func reorderedPowerOf2(n int) bool {
	if 1 == n {
		return true
	}
	s := strconv.Itoa(n)
	dg(s, "", 0, len(s), make([]bool, len(s)))
	return flag
}

var flag bool = false

func dg(n string, re string, index int, len int, record []bool) bool {

	if index == len {
		num, _ := strconv.Atoi(re)
		if isPower(num) {
			flag = true
			return true
		} else {
			return false
		}
	}

	for i := 0; i < len; i++ {

		// 选择
		if !record[i] {
			if n[i] == '0' {
				continue
			}
			record[i] = true
			tmpflag := dg(n, re+string(n[i]), index+1, len, record)
			record[i] = false
			if tmpflag {
				return true
			}

		}
	}

	return false
}

func isPower(n int) bool {

	i := 1
	for i <= n {
		if i == n {
			return true
		} else {
			i = 2 * i
		}
	}
	return false
}

// @lc code=end
