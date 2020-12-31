// 1046. 最后一块石头的重量
// https://leetcode-cn.com/problems/last-stone-weight/

package question1046

import (
	"math"
	"testing"
)

func Test_lastStoneWeight(t *testing.T) {
	// t.Log(lastStoneWeight([]int{2, 7, 4, 1, 8, 1}))

	t.Log(lastStoneWeight([]int{7, 6, 7, 6, 9}))
}

func lastStoneWeight(stones []int) int {
	bigHeap(stones)

	for {
		n1 := pop(&stones)
		n2 := pop(&stones)
		if n1 == 0 || n2 == 0 {
			return n1
		}

		if x := int(math.Abs(float64(n1 - n2))); x > 0 {
			push(&stones, x)
		}
	}
}

func bigHeap(arr []int) {
	// 初始化堆
	for i := len(arr) - 1; i >= 0; i-- {
		down(arr, i, len(arr))
	}
}

func pop(arr *[]int) int {
	n := len(*arr) - 1
	if n < 0 {
		return 0
	}

	x := (*arr)[0]
	(*arr)[0], (*arr)[n] = (*arr)[n], (*arr)[0]
	*arr = (*arr)[:n]
	if n > 0 {
		down(*arr, 0, n)
	}

	return x
}

func push(arr *[]int, x int) {
	*arr = append(*arr, x)
	up(*arr, len(*arr)-1)
}

// 堆节点下沉
func down(arr []int, i, l int) {
	tmp := arr[i]
	for k := i*2 + 1; k < l; k = k*2 + 1 {
		if k+1 < l && arr[k] < arr[k+1] {
			k++
		}

		if tmp < arr[k] {
			arr[i] = arr[k]
			i = k
		} else {
			break
		}
	}

	arr[i] = tmp
}

func up(arr []int, j int) {
	tmp := arr[j]
	for {
		i := (j - 1) / 2 // parent
		if i == j || tmp <= arr[i] {
			break
		}
		arr[j] = arr[i]
		j = i
	}
	arr[j] = tmp
}
