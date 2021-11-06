package questions

import "unicode"

/*
 * @lc app=leetcode.cn id=500 lang=golang
 *
 * [500] 键盘行
 */

// @lc code=start
func findWords(words []string) []string {

	var chacter [26]rune = [26]rune{1, 2, 2, 1, 0, 1, 1, 1, 0, 1, 1, 1, 2, 2, 0, 0, 0, 0, 1, 0, 0, 2, 0, 2, 0, 2}

	resultStr := []string{}

	for _, word := range words {

		var result bool = true
		firstChIdx := chacter[unicode.ToLower(rune(word[0]))-'a']
		for _, ch := range word[1:] {

			currentIdex := chacter[unicode.ToLower(rune(ch))-'a']

			if firstChIdx != currentIdex {

				result = false
				break
			}
		}
		if result {
			resultStr = append(resultStr, word)
		}

	}

	return resultStr
}

// @lc code=end
