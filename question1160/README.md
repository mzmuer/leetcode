## 题目：[拼写单词](https://leetcode-cn.com/problems/find-words-that-can-be-formed-by-characters/)

给你一份『词汇表』（字符串数组） `words` 和一张『字母表』（字符串） `chars`。

假如你可以用 `chars` 中的『字母』（字符）拼写出 `words` 中的某个『单词』（字符串），那么我们就认为你掌握了这个单词。

注意：每次拼写时，`chars` 中的每个字母都只能用一次。

返回词汇表 `words` 中你掌握的所有单词的 **`长度之和`**。

**示例1:**
>输入：words = ["cat","bt","hat","tree"], chars = "atach"  
>输出：6  
>解释：   
>可以形成字符串 "cat" 和 "hat"，所以答案是 3 + 3 = 6。

**示例2:**
>输入：words = ["hello","world","leetcode"], chars = "welldonehoneyr"  
>输出：10  
>解释：  
>可以形成字符串 "hello" 和 "world"，所以答案是 5 + 5 = 10。  

提示：
1. `1 <= words.length <= 1000`
2. `1 <= words[i].length, chars.length <= 100`
3. 所有字符串中都仅包含小写英文字母

## 思路
1. 可以用一个长度26的数组记录`chars`中每个字母的数量。
2. 遍历`words`，如果每个单词字母数量不多余`chars`中的字母就是已经掌握。

## 实现
```go
func countCharacters(words []string, chars string) int {
	alphabet := [26]int{}
	for _, v := range chars {
		alphabet[v-'a']++
	}

	var ans int
	for _, word := range words {
		w := [26]int{}
		for _, v := range word {
			w[v-'a']++
		}

		flag := true
		for i, v := range w {
			if v > alphabet[i] {
				flag = false
			}
		}

		if flag {
			ans += len(word)
		}
	}

	return ans
}
```