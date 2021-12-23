## 题目：[鸡蛋掉落](https://leetcode-cn.com/problems/super-egg-drop/)

你将获得 `K` 个鸡蛋，并可以使用一栋从 `1` 到 `N`  共有 `N` 层楼的建筑。

每个蛋的功能都是一样的，如果一个蛋碎了，你就不能再把它掉下去。

你知道存在楼层 `F` ，满足 `0 <= F <= N` 任何从高于 `F` 的楼层落下的鸡蛋都会碎，从 `F` 楼层或比它低的楼层落下的鸡蛋都不会破。

每次移动，你可以取一个鸡蛋（如果你有完整的鸡蛋）并把它从任一楼层 `X` 扔下（满足 `1 <= X <= N`）。

你的目标是**确切地**知道 `F` 的值是多少。

无论 `F` 的初始值如何，你确定 `F` 的值的最小移动次数是多少？

**示例1:**
>输入：K = 1, N = 2  
输出：2  
解释：  
鸡蛋从 1 楼掉落。如果它碎了，我们肯定知道 F = 0 。  
否则，鸡蛋从 2 楼掉落。如果它碎了，我们肯定知道 F = 1 。  
如果它没碎，那么我们肯定知道 F = 2 。  
因此，在最坏的情况下我们需要移动 2 次以确定 F 是多少。  

**示例2:**
>输入：K = 2, N = 6  
输出：3

**示例3:**
>输入：K = 3, N = 14  
输出：4

提示：
* `1 <= K <= 100`
* `1 <= N <= 10000`

## 思路
1. 动态规划+二分的问题
2. 每次从X楼扔下鸡蛋后会有两种状态
* 鸡蛋没有碎，状态为`(K, N-X)`
* 鸡蛋碎了，状态为`(K-1,X-1)`
3. 所以定义`dp(K,N)`为状态`(K,N)`下的最少步数。然后我们使用二分来确定`X`就好了


## [实现](https://github.com/mzmuer/leetcode/blob/master/question8/answer_test.go)
```go
var memo = map[int]int{}

func superEggDrop(K int, N int) int {
	return dp(K, N)
}

func dp(K, N int) int {
	if _, ok := memo[N*100+K]; !ok {
		var ans int
		if N == 0 {
			ans = 0
		} else if K == 1 {
			ans = N
		} else {
			lo := 1
			hi := N
			for lo+1 < hi {
				x := (lo + hi) / 2
				t1 := dp(K-1, x-1)
				t2 := dp(K, N-x)

				if t1 < t2 {
					lo = x
				} else if t1 > t2 {
					hi = x
				} else {
					lo = x
					hi = x
				}
			}

			ans = 1 + _min(_max(dp(K-1, lo-1), dp(K, N-lo)),
				_max(dp(K-1, hi-1), dp(K, N-hi)))
		}

		memo[N*100+K] = ans
	}

	return memo[N*100+K]
}

func _max(n1, n2 int) int {
	return int(math.Max(float64(n1), float64(n2)))
}

func _min(n1, n2 int) int {
	return int(math.Min(float64(n1), float64(n2)))
}
```