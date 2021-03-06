## 题目：[使数组唯一的最小增量](https://leetcode-cn.com/problems/minimum-increment-to-make-array-unique/)

给定整数数组 `A`，每次 *move* 操作将会选择任意 `A[i]`，并将其递增 `1`。

返回使 `A` 中的每个值都是唯一的最少操作次数。

**示例1:**
>输入：[1,2,2]  
>输出：1  
>解释：经过一次 move 操作，数组将变为 [1, 2, 3]。

**示例2:**
>输入：[3,2,1,2,1,7]  
>输出：6  
>解释：经过 6 次 move 操作，数组将变为 [3, 4, 1, 2, 5, 7]。  
>可以看出 5 次或 5 次以下的 move 操作是不能让数组的每个值唯一的。

提示：
1. `0 <= A.length <= 40000`
2. `0 <= A[i] < 40000`

## 思路
1. 根据提示，可以创建一个`80000`的数组存下最后的所有结果(*`A中有40000个数据且都为39999`*)
2. 先将`A`中的数据填入数组。然后遍历数组
3. 有重复的值就需要向前操作。假如这样一个数组`[1,1,1,1]`,最短不重复应该为`[1,2,3,4]`,也就是是需要将其中三个`1`增加到2,3,4,那么增加的次是也就是`2+3+4-(1+1+1)`。
4. 也就是说有`n`个重复的`x`,需要增加为不重复的`x,x1,x2...xn`,需要递增`(x1+x2+...+xn)-x(n-1)`

## [实现](https://github.com/mzmuer/leetcode/blob/master/question945/answer_test.go)
```go
func minIncrementForUnique(A []int) int {
	arr := [80000]int{}
	for _, v := range A {
		arr[v]++
	}

	var (
		ans   int
		token int
	)
	for i := 0; i < 80000; i++ {
		if arr[i] > 1 {
			token += arr[i] - 1
			ans -= i * (arr[i] - 1) // 减去x(n-1)
		} else if token > 0 && arr[i] == 0 {
			token--
			ans += i // 增加 xn
		}
	}

	return ans
}
```