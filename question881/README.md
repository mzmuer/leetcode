## 题目：[救生艇](https://leetcode-cn.com/problems/boats-to-save-people/)

第 `i` 个人的体重为 `people[i]`，每艘船可以承载的最大重量为 `limit`。  
每艘船最多可同时载两人，但条件是这些人的重量之和最多为 `limit`。  
返回载到每一个人所需的最小船数。(保证每个人都能被船载)。

**示例1:**
>输入：people = [1,2], limit = 3  
 输出：1  
 解释：1 艘船载 (1, 2)  

**示例2:**
>输入：people = [3,2,2,1], limit = 3  
 输出：3  
 解释：3 艘船分别载 (1, 2), (2) 和 (3)  

**示例3:**
>输入：people = [3,5,3,4], limit = 5  
 输出：4  
 解释：4 艘船分别载 (3), (3), (4), (5)

## 思路
1. 贪心,先选择最轻的人  
如果我能和最重的一起走，那就一起走    
如果不能，那么最重的人只能单独走  
剩下的人再循环

## [实现](https://github.com/mzmuer/leetcode/blob/master/question881/answer_test.go)
```go
func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	var ans int
	left, right := 0, len(people)-1

	for left <= right {
		if people[left]+people[right] > limit {
			right--
		} else {
			right--
			left++
		}
		ans++
	}

	return ans
}
```