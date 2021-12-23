## 题目：[快乐数](https://leetcode-cn.com/problems/happy-number/)

编写一个算法来判断一个数 `n` 是不是快乐数。

「快乐数」定义为：对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和，然后重复这个过程直到这个数变为 1，也可能是 **无限循环** 但始终变不到 1。如果 可以变为  1，那么这个数就是快乐数。

如果 n 是快乐数就返回 `True` ；不是，则返回 `False` 。

**示例:**
> 输入：19  
输出：true  
解释：  
1 + 81 = 82  
64 + 4 = 68  
36 + 64 = 100  
1 + 0 + 0 = 1

     
## 思路
1. 获取数组的每一位的平方和，直到等于1，或者出现重复的数字（无限循环）。

## [实现](https://github.com/mzmuer/leetcode/blob/master/question202/answer_test.go)
```go
func isHappy(n int) bool {
	var filter = make(map[int]struct{})

	for {
		if n == 1 {
			return true
		}

		if _, ok := filter[n]; ok {
			return false
		}

		filter[n] = struct{}{}
		tmp := 0
		for n != 0 {
			mod := n % 10
			tmp += mod * mod
			n /= 10
		}
		n = tmp
	}
}
```