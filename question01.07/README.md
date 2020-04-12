## 题目：[旋转矩阵](https://leetcode-cn.com/problems/rotate-matrix-lcci/)

给你一幅由 `N × N` 矩阵表示的图像，其中每个像素的大小为 4 字节。请你设计一种算法，将图像旋转 90 度。

不占用额外内存空间能否做到？

**示例1:**
>给定 matrix =   
>[  
>  [1,2,3],  
>  [4,5,6],  
>  [7,8,9]  
>],
>
>原地旋转输入矩阵，使其变为:  
>[  
>  [7,4,1],  
>  [8,5,2],  
>  [9,6,3]  
>]  

**示例2:**
>给定 matrix =   
>[  
>  [ 5, 1, 9,11],  
>  [ 2, 4, 8,10],  
>  [13, 3, 6, 7],  
>  [15,14,12,16]  
>],   
>
>原地旋转输入矩阵，使其变为:  
>[  
>  [15,13, 2, 5],  
>  [14, 3, 4, 1],  
>  [12, 6, 8, 9],  
>  [16, 7,10,11]  
>]  

## 思路
1. 旋转之后的关系为`matrix[n-1-i][j]=matrix[i][j]`
2. 使用一个辅助矩阵来帮助旋转

## [实现](https://github.com/mzmuer/leetcode/blob/master/question8/answer_test.go)
```go
func rotate(matrix [][]int) {
	copyMatrix := make([][]int, len(matrix))
	for i := range matrix {
		copyMatrix[i] = make([]int, len(matrix[i]))
		copy(copyMatrix[i], matrix[i])
	}

	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[j][n-i-1] = copyMatrix[i][j]
		}
	}
}
```