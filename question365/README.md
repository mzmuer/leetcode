## 题目：[水壶问题](https://leetcode-cn.com/problems/water-and-jug-problem/)

有两个容量分别为 x升 和 y升 的水壶以及无限多的水。请判断能否通过使用这两个水壶，从而可以得到恰好 z升 的水？

如果可以，最后请用以上水壶中的一或两个来盛放取得的 z升 水。

你允许：
* 装满任意一个水壶
* 清空任意一个水壶
* 从一个水壶向另外一个水壶倒水，直到装满或者倒空

**示例1:**
>输入: x = 3, y = 5, z = 4  
>输出: True  

**示例2:**
>输入: x = 2, y = 6, z = 5  
>输出: False
     
## 思路
1. [贝祖定理](https://baike.baidu.com/item/%E8%A3%B4%E8%9C%80%E5%AE%9A%E7%90%86/5186593?fromtitle=%E8%B4%9D%E7%A5%96%E5%AE%9A%E7%90%86&fromid=5185441)

## 实现
```go
func canMeasureWater(x int, y int, z int) bool {
	if x+y < z {
		return false
	}

	if x == 0 || y == 0 {
		return z == 0 || x+y == z
	}

	return z%_gcd(x, y) == 0
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