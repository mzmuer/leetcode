// 22. 括号生成
// https://leetcode-cn.com/problems/generate-parentheses/

package question22

import (
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	t.Log(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	list := make([]string, 0)
	dp(2*n-1, "(", &list)

	ans := make([]string, 0)
	for _, v := range list {
		if valid(v) {
			ans = append(ans, v)
		}
	}

	return ans
}

func valid(s string) bool {
	var balance int
	for _, c := range s {
		if c == '(' {
			balance++
		} else if balance--; balance < 0 {
			return false
		}
	}

	return balance == 0
}

func dp(n int, s string, list *[]string) {
	if n == 0 {
		*list = append(*list, s)
		return
	}

	dp(n-1, s+"(", list)
	dp(n-1, s+")", list)
}
