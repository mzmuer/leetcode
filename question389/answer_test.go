// 389. 找不同
// https://leetcode-cn.com/problems/find-the-difference/

package question389

import (
	"testing"
)

func Test_findTheDifference(t *testing.T) {
	t.Log(findTheDifference("abcd", "abcde"))
}

// func findTheDifference(s string, t string) byte {
// 	var arr [26]int
// 	for _, c := range s {
// 		arr[c-'a']++
// 	}
//
// 	for _, c := range t {
// 		if arr[c-'a']--; arr[c-'a'] < 0 {
// 			return byte(c)
// 		}
// 	}
// 	return 0
// }

func findTheDifference(s, t string) (diff byte) {
	for i := range s {
		diff ^= s[i] ^ t[i]
	}
	return diff ^ t[len(t)-1]
}
