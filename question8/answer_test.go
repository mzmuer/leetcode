// 8. 字符串转换整数 (atoi)
// https://leetcode-cn.com/problems/string-to-integer-atoi/

package question289

import (
	"math"
	"testing"
)

func Test_myAtoi(t *testing.T) {
	t.Log(myAtoi("42") == 42)
}

func myAtoi(str string) int {
	var validNum = make([]byte, 0)
	for _, c := range str {
		if len(validNum) == 0 && c == ' ' {
			continue
		}

		if c != '-' && c != '+' && (c < '0' || c > '9') {
			if len(validNum) == 0 || (len(validNum) == 1 && (validNum[0] == '-' || validNum[0] == '+')) {
				return 0
			}

			break
		}

		if len(validNum) != 0 && (c == '-' || c == '+') {
			if len(validNum) == 1 {
				return 0
			}
			break
		}

		validNum = append(validNum, byte(c))
	}

	if len(validNum) == 0 {
		return 0
	}

	// ---
	neg := false
	if validNum[0] == '+' {
		validNum = validNum[1:]
	} else if validNum[0] == '-' {
		neg = true
		validNum = validNum[1:]
	}

	l := len(validNum)
	var integer float64
	for i, v := range validNum {
		v -= '0'
		integer += float64(v) * math.Pow10(l-i-1)

	}

	if neg {
		integer = -integer
	}

	if integer > math.MaxInt32 {
		return math.MaxInt32
	} else if integer < math.MinInt32 {
		return math.MinInt32
	}

	return int(integer)
}
