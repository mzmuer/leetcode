## 题目：[航班预订统计](https://leetcode-cn.com/problems/corporate-flight-bookings/)

这里有 `n` 个航班，它们分别从 `1` 到 `n` 进行编号。

有一份航班预订表 `bookings` ，表中第 `i` 条预订记录 `bookings[i] = [firsti, lasti, seatsi]` 意味着在从 `firsti` 到 `lasti` （包含 `firsti` 和 `lasti` ）的 **每个航班** 上预订了 `seatsi` 个座位。

请你返回一个长度为 `n` 的数组 `answer`，其中 `answer[i]` 是航班 i 上预订的座位总数。

**示例1:**
>输入：bookings = [[1,2,10],[2,3,20],[2,5,25]], n = 5  
 输出：[10,55,45,25,25]  
 解释：  
 航班编号        1   2   3   4   5  
 预订记录 1 ：   10  10  
 预订记录 2 ：       20  20  
 预订记录 3 ：       25  25  25  25  
 总座位数：      10  55  45  25  25  
 因此，answer = [10,55,45,25,25]

**示例2:**
>输入：bookings = [[1,2,10],[2,2,15]], n = 2  
 输出：[10,25]  
 解释：  
 航班编号        1   2  
 预订记录 1 ：   10  10  
 预订记录 2 ：       15  
 总座位数：      10  25  
 因此，answer = [10,25]  
     
## 思路
1. 使用差分数组，如果对一个数组的区间进行频繁的增减，而读取较少，可以考虑使用差分数组
2. 当我们预定了航班`1~2` 只用修改`d[0]+=10` 和`d[2]-=10`。预定了`2~3`同理，修改`d[1]+=20 d[3]-=20`
3. 重复每一份航班预定表得到最后的差分数组，然后遍历差分数组`d`。根据`arr[i] = d[i]+d[i-1]` 得到最后的数组

## [实现](https://github.com/mzmuer/leetcode/blob/master/question1109/answer_test.go)
```go
func corpFlightBookings(bookings [][]int, n int) []int {
	ans := make([]int, n)
	for _, v := range bookings {
		ans[v[0]-1] += v[2]
		if v[1] < n {
			ans[v[1]] -= v[2]
		}
	}

	for i := 1; i < n; i++ {
		ans[i] += ans[i-1]
	}

	return ans
}
```