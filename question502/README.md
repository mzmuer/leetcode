## 题目：[IPO](https://leetcode-cn.com/problems/ipo/)

假设 力扣（LeetCode）即将开始 **IPO** 。为了以更高的价格将股票卖给风险投资公司，力扣 希望在 IPO 之前开展一些项目以增加其资本。 由于资源有限，它只能在 IPO 之前完成最多 `k` 个不同的项目。帮助 力扣 设计完成最多 `k` 个不同项目后得到最大总资本的方式。

给你 `n` 个项目。对于每个项目 `i` ，它都有一个纯利润 `profits[i]` ，和启动该项目需要的最小资本 `capital[i]` 。

最初，你的资本为 `w` 。当你完成一个项目时，你将获得纯利润，且利润将被添加到你的总资本中。

总而言之，从给定项目中选择 **最多** `k` 个不同项目的列表，以 **最大化最终资本** ，并输出最终可获得的最多资本。

答案保证在 32 位有符号整数范围内。

**示例1:**
>输入：k = 2, w = 0, profits = [1,2,3], capital = [0,1,1]  
 输出：4  
 解释：  
 由于你的初始资本为 0，你仅可以从 0 号项目开始。  
 在完成后，你将获得 1 的利润，你的总资本将变为 1。  
 此时你可以选择开始 1 号或 2 号项目。  
 由于你最多可以选择两个项目，所以你需要完成 2 号项目以获得最大的资本。  
 因此，输出最后最大化的资本，为 0 + 1 + 3 = 4。  

**示例2:**
>输入：k = 3, w = 0, profits = [1,2,3], capital = [0,1,2]  
 输出：6

提示：
* `1 <= k <= 105`
* `0 <= w <= 109`
* `n == profits.length`
* `n == capital.length`
* `1 <= n <= 105`
* `0 <= profits[i] <= 104`
* `0 <= capital[i] <= 109`

## 思路
1. 先按照启动资金把 capital 和 profits 排序（方便获取 capital 小于等于 w 的 profits 数据）
2. 首先把启动资金小于w的纯利放入大根堆。
3. 然后完成大根堆顶，也就是获取最高的利润的任务，w+max(profits)
4. 然后再次把小于w的 capital 对应的 profits 放入堆中
5. 接着重复3，直到达到项目数k。或者堆为空

## 实现
```go
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
```