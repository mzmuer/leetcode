## 题目：[滑动窗口中位数](https://leetcode-cn.com/problems/sliding-window-median/)

中位数是有序序列最中间的那个数。如果序列的长度是偶数，则没有最中间的数；此时中位数是最中间的两个数的平均数。

例如：

`[2,3,4]`，中位数是 `3`  
`[2,3]`，中位数是 `(2 + 3) / 2 = 2.5`  
给你一个数组 *nums*，有一个长度为 *k* 的窗口从最左端滑动到最右端。窗口中有 *k* 个数，每次窗口向右移动 *1* 位。你的任务是找出每次窗口移动后得到的新窗口中元素的中位数，并输出由它们组成的数组。


**示例1:**
>给出 nums = `[1,3,-1,-3,5,3,6,7]`，以及 *k* = 3。
>```jso>
>窗口位置                        中位数
> ---------------               -----
> [1  3  -1] -3  5  3  6  7       1
>  1 [3  -1  -3] 5  3  6  7      -1
>  1  3 [-1  -3  5] 3  6  7      -1
>  1  3  -1 [-3  5  3] 6  7       3
>  1  3  -1  -3 [5  3  6] 7       5
>  1  3  -1  -3  5 [3  6  7]      6
>```
> 因此，返回该滑动窗口的中位数数组 `[1,-1,-1,3,5,6]`。

## 思路
1. 利用两个堆，一个大根堆，一个小根堆，分别存储一半的数据
2. 大根堆存小的一半，大根堆存大的一半，只有通过两个堆的堆顶就可以获得中位数。
3. 滑动窗口时，不断将新元素入堆，旧元素删除。删除通过保存下要删除的数据，在获取堆顶元素时间延迟删除
4. 插入和删除元素后，平衡两个堆的元素数量，确保各保存一半数据

## [实现](https://github.com/mzmuer/leetcode/blob/master/question480/answer_test.go)
```go
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
```