// 8. 字符串转换整数 (atoi)
// https://leetcode-cn.com/problems/string-to-integer-atoi/

package question151

import (
	"strings"
	"testing"
)

func Test_reverseWords(t *testing.T) {
	t.Log(reverseWords("the sky is blue") == "blue is sky the")

	t.Log(reverseWords("  hello world!  ") == "world! hello")

	t.Log(reverseWords("a good   example") == "example good a")
}

// 无空格字符构成一个单词。
// 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
// 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
func reverseWords(s string) string {
	var (
		words []string
		cur   string
	)

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if len(cur) > 0 {
				words = append(words, cur)
			}
			cur = ""
			continue
		}

		cur = string(s[i]) + cur
	}

	if len(cur) > 0 {
		words = append(words, cur)
	}

	return strings.Join(words, " ")
}
