## 题目：[找不同](https://leetcode-cn.com/problems/find-the-difference/)

给定两个字符串 *s* 和 *t*，它们只包含小写字母。

字符串 *t* 由字符串 *s* 随机重排，然后在随机位置添加一个字母。

请找出在 *t* 中被添加的字母。

**示例1:**
>输入：s = "abcd", t = "abcde"  
>输出："e"  
>解释：'e' 是那个被添加的字母。

**示例2:**
>输入：s = "", t = "y"  
>输出："y"

**示例3:**
>输入：s = "a", t = "aa"  
>输出："a"  

**示例4:**
>输入：s = "ae", t = "aea"  
>输出："a"

提示:
* `0 <= s.length <= 1000`
* `t.length == s.length + 1`
* `s 和 t 只包含小写字母`
     
## 思路
1. 计数，通过一个数组，计数s出现的字母的个数。  
然后用t中出现的字母减去计数，最终计数结果为负数的字母就是结果值。

2. 位运算，将两个字符串的值全部异或运算之后，结果就是多出的字符。

## 实现
```go
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
```