## 题目：[验证回文串](https://leetcode-cn.com/problems/valid-palindrome/)

给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

**说明：** 本题中，我们将空字符串定义为有效的回文串。

**示例1:**
>输入: "A man, a plan, a canal: Panama"  
>输出: true  

**示例2:**
>输入: "race a car"  
输出: false  
     
## 思路
1. 双指针遍历。
2. 使用双指针从两头遍历，知道`left<=right`那么该字符串就是一个回文串。
3. 在比较过程中，需要跳过非字母和数字的字符。

## [实现](https://github.com/mzmuer/leetcode/blob/master/question125/answer_test.go)
```go
func isPalindrome(s string) bool {
	if s == "" {
		return true
	}

	s = strings.ToLower(s)
	var (
		left, right = 0, len(s) - 1
		ans         = true
	)

	for left <= right {
		if !_isalnum(s[left]) {
			left++
			continue
		}

		if !_isalnum(s[right]) {
			right--
			continue
		}

		if s[left] != s[right] {
			ans = false
			break
		}

		left++
		right--
	}

	return ans
}

func _isalnum(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
```