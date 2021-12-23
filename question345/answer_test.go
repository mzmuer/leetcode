// 345. 反转字符串中的元音字母
// https://leetcode-cn.com/problems/reverse-vowels-of-a-string/

package question345

import (
	"bytes"
	"testing"
)

func Test_reverseVowels(t *testing.T) {
	t.Log(reverseVowels("hello"))
	t.Log(reverseVowels("leetcode"))
}

func isVowel(b byte) bool {
	return bytes.Contains([]byte("aeiouAEIOU"), []byte{b})
}

func reverseVowels(s string) string {
	left, right := 0, len(s)-1
	t := []byte(s)
	for left < right {
		if !isVowel(s[left]) {
			left++
		} else if !isVowel(s[right]) {
			right--
		} else {
			t[left], t[right] = t[right], t[left]
			left++
			right--
		}
	}

	return string(t)
}
