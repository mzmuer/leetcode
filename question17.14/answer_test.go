// 面试题 17.14. 最小K个数
// https://leetcode-cn.com/problems/smallest-k-lcci/

package question17_14

import (
	"testing"
)

func Test_smallestK(t *testing.T) {
	t.Log(smallestK([]int{1, 2, 3}, 0))

	t.Log(smallestK([]int{1, 3, 5, 7, 2, 4, 6, 8}, 4))
}

func smallestK(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}

	heap := arr[:k]

	// 初始化一个大根堆
	for i := k/2 - 1; i >= 0; i-- {
		down(arr, i, k)
	}

	for i := k; i < len(arr); i++ {
		if heap[0] > arr[i] {
			heap[0] = arr[i]
			down(heap, 0, k)
		}
	}

	return heap
}

// 堆下沉
func down(arr []int, i, l int) {
	tmp := arr[i]
	for k := i*2 + 1; k < l; k = 2*k + 1 {
		// 找到左右子节点中，最大的那个子节点
		if k+1 < l && arr[k] < arr[k+1] {
			k++
		}

		// 如果比子节点小，下沉
		if tmp < arr[k] {
			arr[i] = arr[k]
			i = k
		} else {
			// 如果大于等于只节点，下沉结束
			break
		}
	}

	arr[i] = tmp
}
