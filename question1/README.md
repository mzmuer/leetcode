## 题目：[两数之和](https://leetcode-cn.com/problems/two-sum/)

给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

**示例:**
> 给定 nums = [2, 7, 11, 15], target = 9
>
>因为 nums[0] + nums[1] = 2 + 7 = 9
>所以返回 [0, 1]


## 思路
利用map存储已经出现过的数字，遍历nums元素，计算target差值。然后在map中查找

## 实现
```go
func twoSum(nums []int, target int) []int {
    m := map[int]int{}
    for i,v:=range nums{
        diff := target-v
        if index,ok:= m[diff];ok{
            return []int{index,i}
        }

        m[v]=i
    }

    return []int{}
}
```