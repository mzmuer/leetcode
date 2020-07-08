// 面试题40. 最小的k个数
// https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/

package question_40

import (
	"testing"
)

func Test_getLeastNumbers(t *testing.T) {
	a1 := []int{5, 7, 2, 1, 3, 4}
	t.Log(getLeastNumbers(a1, 2))

	// str1 = "ABABAB"
	// str2 = "ABAB"
	// t.Log(gcdOfStrings(str1, str2) == "AB")
	//
	// str1 = "LEET"
	// str2 = "CODE"
	// t.Log(gcdOfStrings(str1, str2) == "")
}

func getLeastNumbers(arr []int, k int) []int {
	sortHeap(arr)
	return arr[:k]

	// var ans = arr[:k]
	// for i := len(ans)/2 - 1; i >= 0; i-- {
	//	adjustHeap(ans, i, len(ans))
	// }
	//
	// for i := k; i < len(arr); i++ {
	//	if ans[0] > arr[i] {
	//		ans[0] = arr[i]
	//		adjustHeap(ans, 0, len(ans))
	//	}
	// }
	//
	// return ans
}

// heap
func sortHeap(arr []int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		adjustHeap(arr, i, len(arr))
	}

	for j := len(arr) - 1; j > 0; j-- {
		arr[0], arr[j] = arr[j], arr[0]
		adjustHeap(arr, 0, j) // 重新对堆进行调整
	}
}

func adjustHeap(arr []int, i, l int) {
	tmp := arr[i]
	for k := i*2 + 1; k < l; k = k*2 + 1 {
		if k+1 < l && arr[k] < arr[k+1] {
			k++
		}

		if arr[k] > tmp {
			arr[i] = arr[k]
			i = k
		} else {
			break
		}
	}

	arr[i] = tmp
}
