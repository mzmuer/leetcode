## 题目：[最长回文串](https://leetcode-cn.com/problems/longest-palindrome/)

给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。

在构造过程中，请注意区分大小写。比如 "Aa" 不能当做一个回文字符串。

注意:
假设字符串的长度不会超过 1010。

**示例1:**
>输入:  
>"abccccdd"  
>
>输出:  
>7  
>  
>解释:  
>我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。

## 思路
1. 用一个数组存储每个字母的数量。
2. 为偶数个的字母可以全部作为回文。
3. 有奇数个的字母可以使用 `n-1(n/2*2)`个做为回文。
4. 可以选择一个奇数长度的的字母作为回文的中心，且只能选择一个。

## 实现
```go
func longestPalindrome(s string) int {
	chars := [52]int{}
	for _, c := range s {
		if c >= 'a' {
			chars[25+c-'a']++
		} else {
			chars[c-'A']++
		}
	}

	var ans int
	for _, v := range chars {
		if v%2 == 0 {
			ans += v
		} else {
			if ans%2 == 0 {
				ans += v
			} else {
				ans += v / 2 * 2
			}
		}
	}

	return ans
}
```