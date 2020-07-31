## 题目：[判断子序列](https://leetcode-cn.com/problems/is-subsequence/)

给定字符串 **s** 和 **t** ，判断 **s** 是否为 **t** 的子序列。

你可以认为 **s** 和 **t** 中仅包含英文小写字母。字符串 **t** 可能会很长（长度 ~= 500,000），而 **s** 是个短字符串（长度 <=100）。

字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，`"ace"`是`"abcde"`的一个子序列，而`"aec"`不是）。

示例 1:  
>**s** = `"abc"`, **t** = `"ahbgdc"`  
>返回 true.

示例 2:  
>**s** = `"axc"`, **t** = `"ahbgdc"`  
返回 `false`.

     
## 思路
1. 只要找到任意一种`s`在`t`中出现的方式就可以了,所以只用贪心的去匹配最考前的值，就是最优的策略。
2. 遍历`t`,如果有和s中匹配的字母，那么`s`的`index`向前移动。直到`index`移动到`s`的结尾。那么`s`就是`t`的子数组。

## [实现](https://github.com/mzmuer/leetcode/blob/master/question392/answer_test.go)
```go
func isSubsequence(s string, t string) bool {
	var (
		l   = len(s)
		idx = 0
	)

	for i := 0; i < len(t) && idx < l; i++ {
		if t[i] == s[idx] {
			idx++
		}
	}

	return idx == l
}
```