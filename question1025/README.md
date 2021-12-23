## 题目：[除数博弈](https://leetcode-cn.com/problems/divisor-game/)

爱丽丝和鲍勃一起玩游戏，他们轮流行动。爱丽丝先手开局。

最初，黑板上有一个数字 N 。在每个玩家的回合，玩家需要执行以下操作：

选出任一 x，满足 0 < x < N 且 N % x == 0 。
用 N - x 替换黑板上的数字 N 。
如果玩家无法执行这些操作，就会输掉游戏。

只有在爱丽丝在游戏中取得胜利时才返回 True，否则返回 false。假设两个玩家都以最佳状态参与游戏。

**示例1:**
>输入：2  
>输出：true  
>解释：爱丽丝选择 1，鲍勃无法进行操作。

**示例2:**
>输入：3  
>输出：false  
>解释：爱丽丝选择 1，鲍勃也选择 1，然后爱丽丝无法进行操作。  
     
## 思路
1. 找规律题，如果`N`为偶数时，先手必胜。
2. `0<x<N`，所以当N为奇数时使得`N%x==0`,那么`x`一定是奇数。奇数减去奇数又是一个偶数。所以偶数先手必胜。

## [实现](https://github.com/mzmuer/leetcode/blob/master/question1025/answer_test.go)
```go
func divisorGame(N int) bool {
	return N%2 == 0
}
```