// 480. 滑动窗口中位数
// https://leetcode-cn.com/problems/sliding-window-median/

package question480

import (
	"container/heap"
	"sort"
	"testing"
)

func Test_medianSlidingWindow(t *testing.T) {
	a1 := []int{1, 3, -1, -3, 5, 3, 6, 7}
	t.Log(medianSlidingWindow(a1, 3))

	a2 := []int{1, 1, 1, 1}
	t.Log(medianSlidingWindow(a2, 2))
}

type xhp struct {
	sort.IntSlice
	size int
}

func (h *xhp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *xhp) Pop() interface{} {
	// a := h.IntSlice
	// v := a[len(a)-1]
	// h.IntSlice = a[:len(a)-1]
	// return v

	v := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return v
}

func (h *xhp) push(v interface{}) {
	h.size++
	heap.Push(h, v)
}

func (h *xhp) pop() interface{} {
	h.size--
	return heap.Pop(h)
}

func (h *xhp) prune() {
	for h.Len() > 0 {
		num := h.IntSlice[0]
		if h == largeHeap {
			num = -num
		}

		if count, has := delayed[num]; has {
			if count > 1 {
				delayed[num]--
			} else {
				delete(delayed, num)
			}
			heap.Pop(h)
		} else {
			break
		}
	}
}

var (
	smallHeap, largeHeap *xhp
	delayed              map[int]int
)

func medianSlidingWindow(nums []int, k int) []float64 {
	delayed = map[int]int{}
	smallHeap = &xhp{} // 小根堆，存数据比较大的那一半
	largeHeap = &xhp{} // 大根堆，存数据比较小的那一半

	mkBalance := func() {
		// 大根堆，比小根堆数量多，大根堆的放入小根堆
		// 小根堆，比大根堆数量多2，小根堆放入大根堆
		if largeHeap.size > smallHeap.size+1 {
			smallHeap.push(-largeHeap.pop().(int))
			largeHeap.prune()
		} else if largeHeap.size < smallHeap.size {
			largeHeap.push(-smallHeap.pop().(int))
			smallHeap.prune()
		}
	}

	insert := func(num int) {
		if largeHeap.Len() == 0 || num < -largeHeap.IntSlice[0] {
			largeHeap.push(-num)
		} else {
			smallHeap.push(num)
		}
		mkBalance()
	}
	erase := func(num int) {
		delayed[num]++
		if -largeHeap.IntSlice[0] >= num {
			largeHeap.size--
			if -largeHeap.IntSlice[0] == num {
				largeHeap.prune()
			}
		} else {
			smallHeap.size--
			if smallHeap.IntSlice[0] == num {
				smallHeap.prune()
			}
		}
		mkBalance()
	}
	getMedian := func() float64 {
		if k&1 > 0 {
			return float64(-largeHeap.IntSlice[0])
		}
		return float64(-largeHeap.IntSlice[0]+smallHeap.IntSlice[0]) / 2
	}

	for i := 0; i < k; i++ {
		insert(nums[i])
	}

	ans := []float64{getMedian()}
	for i := k; i < len(nums); i++ {
		insert(nums[i])
		erase(nums[i-k])
		ans = append(ans, getMedian())
	}

	return ans
}
