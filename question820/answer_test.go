// 820. 单词的压缩编码
// https://leetcode-cn.com/problems/short-encoding-of-words/

package question820

import (
	"testing"
)

func Test_minimumLengthEncoding(t *testing.T) {
	words1 := []string{"time", "me", "bell"}
	t.Log(minimumLengthEncoding(words1) == 10)
}

//1 <= words.length <= 2000
//1 <= words[i].length <= 7
func minimumLengthEncoding(words []string) int {
	var (
		suffix = make(map[string]struct{})
		ans    int
	)

	for _, v := range words {
		suffix[v] = struct{}{}
	}

	for _, v := range words {
		for i := 1; i < len(v); i++ {
			delete(suffix, v[i:])
		}
	}

	for k := range suffix {
		ans += len(k) + 1
	}

	return ans
}
