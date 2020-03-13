## 题目：[多数元素](https://leetcode-cn.com/problems/majority-element/)

给定一个大小为 *n* 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 `⌊ n/2 ⌋` 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。

**示例1:**
>输入: [3,2,3]  
>输出: 3

**示例2:**
>输入: [2,2,1,1,1,2,2]  
>输出: 2
     
## 思路一，使用map
1. 用一个map记录下出现的数字的次数
2. 出现次数最多的就是多数元素

## 实现
```go
func majorityElement(nums []int) int {
	var (
		count         = map[int]int{}
		majority, cnt int
		l             = len(nums)
	)

	for _, num := range nums {
		if count[num]++; count[num] > cnt {
			majority = num
			cnt = count[num]

			// 如果已经是多数了 提前跳出
			if count[num] > l/2 {
				break
			}
		}
	}

	return majority
}
```