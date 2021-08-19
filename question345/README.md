## 题目：[反转字符串中的元音字母](https://leetcode-cn.com/problems/reverse-vowels-of-a-string/)

给你一个字符串 s ，仅反转字符串中的所有元音字母，并返回结果字符串。  

元音字母包括 `'a'`、`'e'`、`'i'`、`'o'`、`'u'`，且可能以大小写两种形式出现。

示例 1:  
>输入：s = "hello"  
输出："holle"

示例 2:  
>输入：s = "leetcode"  
输出："leotcede"

提示： 
* `1 <= s.length <= 3 * 105`
* `s` 由 *可打印的 ASCII* 字符组成
     
## 思路
使用双指针，left和right都获取到元音字母后交换
知道left和right相遇后结束

## [实现](https://github.com/mzmuer/leetcode/blob/master/question345/answer_test.go)
```go
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
```