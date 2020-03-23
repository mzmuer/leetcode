## 题目：[矩形重叠](https://leetcode-cn.com/problems/rectangle-overlap/)

矩形以列表 `[x1, y1, x2, y2]` 的形式表示，其中 `(x1, y1)` 为左下角的坐标，`(x2, y2)` 是右上角的坐标。

如果相交的面积为正，则称两矩形重叠。需要明确的是，只在角或边接触的两个矩形不构成重叠。

给出两个矩形，判断它们是否重叠并返回结果。

**示例1:**
>输入：rec1 = [0,0,2,2], rec2 = [1,1,3,3]  
>输出：true

**示例2:**
>输入：rec1 = [0,0,1,1], rec2 = [1,0,2,1]  
>输出：false

## 思路
如果`rec1`和`rec2`没有重叠，那么只有四种情况，`rec1`在`rec2`的上下左右四个方向。
* `rec1[2] <= rec2[0]` `rec1` 在 `rec2` 的左边
* `rec1[3] <= rec2[1]` `rec1` 在 `rec2` 的下面
* `rec1[0] >= rec2[2]` `rec1` 在 `rec2` 的右边
* `rec1[1] >= rec2[3]` `rec1` 在 `rec2` 的上边

## 实现
```go
func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	return !(rec1[2] <= rec2[0] || // left
		rec1[3] <= rec2[1] || // bottom
		rec1[0] >= rec2[2] || // right
		rec1[1] >= rec2[3]) // top
}
```