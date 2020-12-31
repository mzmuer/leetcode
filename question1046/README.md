## 题目：[最后一块石头的重量](https://leetcode-cn.com/problems/last-stone-weight/)

有一堆石头，每块石头的重量都是正整数。

每一回合，从中选出两块 最重的 石头，然后将它们一起粉碎。假设石头的重量分别为 `x` 和 `y`，且 `x <= y`。那么粉碎的可能结果如下：

* 如果 `x == y`，那么两块石头都会被完全粉碎；
* 如果 `x != y`，那么重量为 x 的石头将会完全粉碎，而重量为 `y` 的石头新重量为 `y-x`。  
最后，最多只会剩下一块石头。返回此石头的重量。如果没有石头剩下，就返回 `0`。

**示例1:**
>输入：[2,7,4,1,8,1]  
>输出：1  
>解释：  
>先选出 7 和 8，得到 1，所以数组转换为 [2,4,1,1,1]，  
>再选出 2 和 4，得到 2，所以数组转换为 [2,1,1,1]，  
>接着是 2 和 1，得到 1，所以数组转换为 [1,1,1]，  
>最后选出 1 和 1，得到 0，最终数组转换为 [1]，这就是最后剩下那块石头的重量。 

提示：
1. `1 <= stones.length <= 30`
2. `1 <= stones[i] <= 1000`
     
## 思路
1. 构造一个大根堆，每次取出最大的两个
2. 粉碎之后如果还有剩余就继续就入堆
3. 直到最后只剩下一个石头，或者一个都不存在，得到最后的结果

## [实现](https://github.com/mzmuer/leetcode/blob/master/question1046/answer_test.go)
```go
func lastStoneWeight(stones []int) int {
	bigHeap(stones)

	for {
		n1, n2 := pop(&stones), pop(&stones)
		if n1 == 0 || n2 == 0 {
			return n1
		}

		if x := n1 - n2; x > 0 {
			push(&stones, x)
		}
	}
}

// 初始化为大根堆
func bigHeap(arr []int) {
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

// 堆节点上升
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
```