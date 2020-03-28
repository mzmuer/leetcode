## 题目：[卡牌分组](https://leetcode-cn.com/problems/x-of-a-kind-in-a-deck-of-cards/)

给定一副牌，每张牌上都写着一个整数。

此时，你需要选定一个数字 `X`，使我们可以将整副牌按下述规则分成 1 组或更多组：

* 每组都有 X 张牌。
* 组内所有的牌上都写着相同的整数。

仅当你可选的 `X >= 2` 时返回 `true`。

**示例1:**
>输入：[1,2,3,4,4,3,2,1]  
>输出：true  
>解释：可行的分组是 [1,1]，[2,2]，[3,3]，[4,4]  

**示例2:**
>输入：[1,1,1,2,2,2,3,3]  
>输出：false  
>解释：没有满足要求的分组。  

**示例3:**
>输入：[1]  
>输出：false  
>解释：没有满足要求的分组。  

**示例4:**
>输入：[1,1]  
>输出：true  
>解释：可行的分组是 [1,1]

**示例5:**
>输入：[1,1,2,2,2,2]  
>输出：true  
>解释：可行的分组是 [1,1]，[2,2]，[2,2]

提示：
1. `1 <= deck.length <= 10000`
2. `0 <= deck[i] < 10000`

## 思路
1. 最大公约数，能分组的的牌一定是每个数字数量的最大公约数。

## [实现](https://github.com/mzmuer/leetcode/blob/master/question914/answer_test.go)
```go
func hasGroupsSizeX(deck []int) bool {
	var count [10000]int
	for _, v := range deck {
		count[v]++
	}

	g := -1
	for i := 0; i < 10000; i++ {
		if count[i] != 0 {
			if g == -1 {
				g = count[i]
			} else {
				g = _gcd(g, count[i])
			}
		}
	}

	return g >= 2
}

func _gcd(x, y int) int {
	tmp := x % y
	if tmp > 0 {
		return _gcd(y, tmp)
	} else {
		return y
	}
}
```