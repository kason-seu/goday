package problems

import (
	"unicode"
	"unicode/utf8"
)

func Q43Reverse(ptr *[5]int) *[5]int {

	for i, j := 0, len(*ptr)-1; i < j; i, j = i+1, j-1 {
		(*ptr)[i], (*ptr)[j] = (*ptr)[j], (*ptr)[i]
	}
	return ptr
}

func Q45RemoveSameLetters(str []string) (re []string) {

	for i := 0; i < len(str)-1; i++ {
		if str[i] != str[i+1] {
			re = append(re, str[i])

			if i == len(str)-2 {
				re = append(re, str[i+1])
			}
		}
	}
	return re
}

func Q46RemoveAdjWhiteSpace(s []byte) []byte {
	isFirstSeriesSpace := false
	i, j := 0, 0
	for i < len(s) {
		// 解析出第一个字符
		r, size := utf8.DecodeRune(s[i:])
		if r == utf8.RuneError {
			panic("runeError")
		}
		// 是空格
		if unicode.IsSpace(r) {
			// 遇到第一个空格开始处理
			if !isFirstSeriesSpace {
				s[j] = ' '
				j++
			}
			isFirstSeriesSpace = true
			i += size
		} else {
			isFirstSeriesSpace = false
			copy(s[j:j+size], s[i:i+size])
			j += size
			i += size
		}

	}

	return s[:j]
}
func Q44Rotate(s []int, left int) []int {

	reverse(s[:left])

	reverse(s[left:])

	reverse(s)

	return s
}
func Q44Rotate2(s []int, left int) []int {

	newS := make([]int, len(s))

	copy(newS[0:left], s[:left])

	copy(s[0:len(s)-left], s[left:])

	copy(s[len(s)-left:], newS[:left])
	return s

}

func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {

		s[i], s[j] = s[j], s[i]
	}
	return s
}

/**
递归求解
 */
func Q47Reverse(s []byte) []byte {
	// 递归终止
	if len(s) == 0 {
		return s
	}
	r, size := utf8.DecodeRune(s)
	if r == utf8.RuneError {
		return nil
	}
	rotate(s, size)
	Q47Reverse(s[0 : len(s)-size])

	return s
}

func rotate(s []byte, left int) []byte {

	reverse2(s[:left])

	reverse2(s[left:])

	reverse2(s)

	return s
}
func reverse2(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {

		s[i], s[j] = s[j], s[i]
	}
	return s
}
