// 1160. 拼写单词
// https://leetcode-cn.com/problems/find-words-that-can-be-formed-by-characters/

package question1160

import (
	"testing"
)

func Test_countCharacters(t *testing.T) {
	words1 := []string{"cat", "bt", "hat", "tree"}
	chars1 := "atach"
	t.Log(countCharacters(words1, chars1) == 6)

	//str1 = "ABABAB"
	//str2 = "ABAB"
	//t.Log(gcdOfStrings(str1, str2) == "AB")
	//
	//str1 = "LEET"
	//str2 = "CODE"
	//t.Log(gcdOfStrings(str1, str2) == "")
}

//1 <= words.length <= 1000
//1 <= words[i].length, chars.length <= 100
//所有字符串中都仅包含小写英文字母
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
