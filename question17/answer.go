// 17. 电话号码的字母组合
// https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/

package question17

var m = map[uint8]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

// 思路：循环调用
func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	return f(digits)
}

func f(digits string) []string {
	var res []string
	if len(digits) == 1 {
		s := m[digits[0]]
		for _, v := range s {
			res = append(res, string(v))
		}
		return res
	} else {
		r := f(digits[1:])
		if s, ok := m[digits[0]]; ok {
			for _, v := range s {
				for _, val := range r {
					res = append(res, string(v)+val)
				}
			}
		} else {
			return r
		}

		return res
	}
}
