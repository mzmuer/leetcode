// 20. 有效的括号
// https://leetcode-cn.com/problems/valid-parentheses/

package question20

// 思路一: 栈
func IsValid(s string) bool {
	stack := make([]int32, 0)
	for _, v := range s {
		if v == '(' || v == '{' || v == '[' {
			stack = append(stack, v)
		} else if v == ')' || v == '}' || v == ']' {
			if len(stack) < 1 {
				return false
			}

			prev := stack[len(stack)-1:]
			if v == ')' {
				if prev[0] != '(' {
					return false
				}
			} else if v == '}' {
				if prev[0] != '{' {
					return false
				}
			} else if v == ']' {
				if prev[0] != '[' {
					return false
				}
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
