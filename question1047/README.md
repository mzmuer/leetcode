## 题目：[删除字符串中的所有相邻重复项](https://leetcode-cn.com/problems/remove-all-adjacent-duplicates-in-string/)

给出由小写字母组成的字符串 `S`，**重复项删除操作**会选择两个相邻且相同的字母，并删除它们。

在 S 上反复执行重复项删除操作，直到无法继续删除。

在完成所有重复项删除操作后返回最终的字符串。答案保证唯一。

**示例1:**
>输入："abbaca"
>输出："ca"  
>解释：  
>例如，在 "abbaca" 中，我们可以删除 "bb" 由于两字母相邻且相同，这是此时唯一可以执行删除操作的重复项。之后我们得到字符串 "aaca"，其中又只有 "aa" 可以执行重复项删除操作，所以最后的字符串为 "ca"。

提示：
1. `1 <= S.length <= 20000`
2. `S 仅由小写英文字母组成。`
     
## 思路
1. 类似于匹配括号，使用栈。
2. 遍历 `S` 每个字符逐步入栈，当元素与栈顶元素相同时，移除栈顶元素，否者入栈。
3. 最后栈中剩下的元素全部弹出，组成的字符串就是结果

## [实现](https://github.com/mzmuer/leetcode/blob/master/question1047/answer_test.go)
```go
func removeDuplicates(S string) string {
	stack := make([]byte, 0)
	for i := range S {
		if len(stack) == 0 || stack[len(stack)-1] != S[i] {
			stack = append(stack, S[i])
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return string(stack)
}
```