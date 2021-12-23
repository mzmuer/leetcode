// 502. IPO
// https://leetcode-cn.com/problems/ipo/

package question502

import (
	"sort"
	"testing"
)

func Test_findMaximizedCapital(t *testing.T) {
	t.Log(findMaximizedCapital(10, 0, []int{1, 2, 3}, []int{0, 1, 1}))
}

// 1. 先按照启动资金把 capital 和 profits 排序（方便获取 capital 小于等于 w 的 profits 数据）
// 2. 首先把启动资金小于w的纯利放入大根堆。
// 3. 然后完成大根堆顶，也就是获取最高的利润的任务，w+max(profits)
// 4. 然后再次把小于w的 capital 对应的 profits 放入堆中
// 5. 接着重复3，直到达到项目数k。或者堆为空
func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	n := len(profits)
	type pair struct{ c, p int }
	arr := make([]pair, n)
	for i, p := range profits {
		arr[i] = pair{capital[i], p}
	}
	sort.Slice(arr, func(i, j int) bool { return arr[i].c < arr[j].c })

	hp := heap{}
	for i, j := 0, 0; i < k; i++ {
		// 初始入堆
		for ; j < n && w >= arr[j].c; j++ {
			hp.push(arr[j].p)
		}
		if hp.len() == 0 {
			break
		}
		w += hp.pop()
	}
	return w
}

type heap struct {
	arr []int
}

func (s heap) len() int {
	return len(s.arr)
}

// 入堆
func (s *heap) push(x int) {
	s.arr = append(s.arr, x)
	s._up(len(s.arr) - 1)
}

// 出堆
func (s *heap) pop() int {
	n := len(s.arr) - 1
	if n < 0 {
		return 0
	}

	top := s.arr[0]
	if n > 0 {
		s.arr[0], s.arr[n] = s.arr[n], s.arr[0]
		s._down(0, n)
	}

	s.arr = s.arr[:n]
	return top
}

// 上升
func (s *heap) _up(i int) {
	tmp := s.arr[i]
	for k := (i - 1) / 2; k != i; k = (k - 1) / 2 {
		if s.arr[k] < tmp {
			s.arr[i] = s.arr[k]
			i = k
		} else {
			break
		}
	}

	s.arr[i] = tmp
}

// 下沉
func (s *heap) _down(i, l int) {
	tmp := s.arr[i]
	for k := 2*i + 1; k < l; k = 2*k + 1 {
		if k+1 < l && s.arr[k+1] > s.arr[k] {
			k++
		}

		if s.arr[k] > tmp {
			s.arr[i] = s.arr[k]
			i = k
		} else {
			break
		}
	}

	s.arr[i] = tmp
}
