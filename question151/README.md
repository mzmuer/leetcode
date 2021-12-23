## 题目：[翻转字符串里的单词](https://leetcode-cn.com/problems/reverse-words-in-a-string/)

给定一个字符串，逐个翻转字符串中的每个单词。

**示例1:**
>输入: "the sky is blue"  
输出: "blue is sky the"

**示例2:**
>输入: "  hello world!  "  
输出: "world! hello"  
解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。

**示例3:**
>输入: "a good   example"  
输出: "example good a"  
解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。

**说明：**

* 无空格字符构成一个单词。
* 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
* 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。

## 思路
1. 从后往前遍历得到每个单词
2. 拼接成字符串返回

## [实现](https://github.com/mzmuer/leetcode/blob/master/question151/answer_test.go)
```go
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
```