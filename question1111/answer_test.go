// 1111. 有效括号的嵌套深度
// https://leetcode-cn.com/problems/maximum-nesting-depth-of-two-valid-parentheses-strings/

package question1111

import (
	"testing"
)

func Test_maxDepthAfterSplit(t *testing.T) {
	// [0,1,1,1,1,0]
	seq1 := "(()())"
	t.Log(maxDepthAfterSplit(seq1))

	// [0,0,0,1,1,0,1,1]
	seq2 := "()(())()"
	t.Log(maxDepthAfterSplit(seq2))
}

// 1 <= text.size <= 10000
func maxDepthAfterSplit(seq string) []int {
	var d int
	var ans []int
	for _, c := range seq {
		if c == '(' {
			d++
			ans = append(ans, d%2)
		} else {
			ans = append(ans, d%2)
			d--
		}
	}
	return ans
}
