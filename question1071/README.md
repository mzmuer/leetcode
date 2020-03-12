## 题目：[字符串的最大公因子](https://leetcode-cn.com/problems/greatest-common-divisor-of-strings/)

对于字符串 `S` 和 `T`，只有在 `S = T + ... + T`（`T` 与自身连接 1 次或多次）时，我们才认定 “`T` 能除尽 `S`”。

返回最长字符串 `X`，要求满足 `X` 能除尽 `str1` 且 `X` 能除尽 `str2`。

**示例1:**
>输入：str1 = "ABCABC", str2 = "ABC"  
>输出："ABC"

**示例2:**
>输入：str1 = "ABABAB", str2 = "ABAB"  
>输出："AB"

**示例3:**
>输入：str1 = "LEET", str2 = "CODE"  
输出：""
     
## 思路
1. `X`一定是str1和str2的公共前缀
2. `X`的长度肯定是len(str1)和len(str2)的最大公约数
3. 判断最大公约数长度的前缀是否能组成str1和str2

## 实现
```go
func gcdOfStrings(str1 string, str2 string) string {
	T := str1[:_gcd(len(str1), len(str2))]
	if check(T, str1) && check(T, str2) {
		return T
	}

	return ""
}

func check(t, s string) bool {
	lenx := len(s) / len(t)
	var ans string
	for i := 1; i <= lenx; i++ {
		ans = ans + t
	}

	return ans == s
}

func _gcd(x, y int) int {
	tmp := x % y
	if tmp > 0 {
		return _gcd(y, tmp)
	} else {
		return y
	}
}
```