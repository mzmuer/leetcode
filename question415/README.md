## 题目：[字符串相加](https://leetcode-cn.com/problems/add-strings/)

给定两个字符串形式的非负整数`num1`和`num2` ，计算它们的和。

**注意**：
1. `num1`和`num2` 的长度都小于`5100`.
2. `num1`和`num2` 都只包含数字`0-9`.
3. `num1`和`num2` 都不包含任何前导零。
4. **你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式**。
     
## 思路
1. 从低位开始遍历相加，注意进位就好

## [实现](https://github.com/mzmuer/leetcode/blob/master/question415/answer_test.go)
```go
func addStrings(num1 string, num2 string) string {
	var ans string
	var d int
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		var x int
		if i >= 0 {
			x = int(num1[i] - '0')
		}

		var y int
		if j >= 0 {
			y = int(num2[j] - '0')
		}

		ans = string((x+y+d)%10+'0') + ans
		d = (x + y + d) / 10
	}

	if d != 0 {
		ans = string('0'+d) + ans
	}

	return ans
}
```