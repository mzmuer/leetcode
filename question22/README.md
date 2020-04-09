## 题目：[括号生成](https://leetcode-cn.com/problems/generate-parentheses/)

数字 *n* 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 **有效的** 括号组合。

**示例1:**
>输入：n = 3  
>输出：[  
       "((()))",  
       "(()())",  
       "(())()",  
       "()(())",  
       "()()()"  
     ]  

## 思路
1. 暴力解法，递归出全部的括号组合
2. 验证所以有效的括号组合

## [实现](https://github.com/mzmuer/leetcode/blob/master/question22/answer_test.go)
```go
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
```