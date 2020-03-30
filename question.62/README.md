## 题目：[圆圈中最后剩下的数字](https://leetcode-cn.com/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/)

0,1,,n-1这n个数字排成一个圆圈，从数字0开始，每次从这个圆圈里删除第m个数字。求出这个圆圈里剩下的最后一个数字。

例如，0、1、2、3、4这5个数字组成一个圆圈，从数字0开始每次删除第3个数字，则删除的前4个数字依次是2、0、4、1，因此最后剩下的数字是3。

**示例1:**
>输入: n = 5, m = 3  
>输出: 3

**示例2:**
>输入: n = 10, m = 17  
>输出: 2

提示：
1. `1 <= n <= 10^5`
2. `1 <= m <= 10^6`

## 思路
1. 这是一个约瑟夫环的问题。

## [实现](https://github.com/mzmuer/leetcode/blob/master/question。62/answer_test.go)
```go
func lastRemaining(n int, m int) int {
	var f int
	for i := 2; i != n+1; i++ {
		f = (m + f) % i
	}
	return f
}
```