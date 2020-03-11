// 1013. 将数组分成和相等的三个部分
// https://leetcode-cn.com/problems/partition-array-into-three-parts-with-equal-sum/

package question1013

import (
	"testing"
)

func Test_canThreePartsEqualSum(t *testing.T) {
	a1 := []int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1}
	t.Log(canThreePartsEqualSum(a1) == true)

	a2 := []int{0, 2, 1, -6, 6, 7, 9, -1, 2, 0, 1}
	t.Log(canThreePartsEqualSum(a2) == false)

	a3 := []int{10, -10, 10, -10, 10, -10, 10, -10}
	t.Log(canThreePartsEqualSum(a3) == true)
}

func canThreePartsEqualSum(A []int) bool {
	if len(A) < 3 {
		return false
	}

	var sum int
	for _, v := range A {
		sum += v
	}

	if sum%3 != 0 {
		return false
	}

	var (
		target = sum / 3
		tmpSum int
	)

	for i := 0; i < len(A)-2; i++ {
		tmpSum += A[i]
		if tmpSum != target {
			continue
		}

		tmpSum = 0
		for j := i + 1; j < len(A)-1; j++ {
			tmpSum += A[j]
			if tmpSum != target {
				continue
			}

			return true
		}

		break
	}

	return false
}
