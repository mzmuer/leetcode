## 题目：[单词的压缩编码](https://leetcode-cn.com/problems/short-encoding-of-words/)

给定一个单词列表，我们将这个列表编码成一个索引字符串 `S` 与一个索引列表 `A`。

例如，如果这个列表是 `["time", "me", "bell"]`，我们就可以将其表示为 `S = "time#bell#"` 和 `indexes = [0, 2, 5]`。

对于每一个索引，我们可以通过从字符串 `S` 中索引的位置开始读取字符串，直到 "#" 结束，来恢复我们之前的单词列表。

那么成功对给定单词列表进行编码的最小字符串长度是多少呢？

**示例1:**
>输入: words = ["time", "me", "bell"]  
>输出: 10  
>说明: S = "time#bell#" ， indexes = [0, 2, 5]  

提示：
1. `<= words.length <= 2000`
2. `1 <= words[i].length <= 7`
3. 每个单词都是小写字母

## 思路
1. 根据提示，可以创建一个`80000`的数组存下最后的所有结果(*`A中有40000个数据且都为39999`*)
2. 先将`A`中的数据填入数组。然后遍历数组
3. 有重复的值就需要向前操作。假如这样一个数组`[1,1,1,1]`,最短不重复应该为`[1,2,3,4]`,也就是是需要将其中三个`1`增加到2,3,4,那么增加的次是也就是`2+3+4-(1+1+1)`。
4. 也就是说有`n`个重复的`x`,需要增加为不重复的`x,x1,x2...xn`,需要递增`(x1+x2+...+xn)-x(n-1)`

## [实现](https://github.com/mzmuer/leetcode/blob/master/question820/answer_test.go)
```go
func minimumLengthEncoding(words []string) int {
	var (
		suffix = make(map[string]struct{})
		ans    int
	)

	for _, v := range words {
		suffix[v] = struct{}{}
	}

	for _, v := range words {
		for i := 1; i < len(v); i++ {
			delete(suffix, v[i:])
		}
	}

	for k := range suffix {
		ans += len(k) + 1
	}

	return ans
}
```