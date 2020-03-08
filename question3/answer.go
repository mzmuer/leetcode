// 3. 无重复字符的最长子串
// https://www.jianshu.com/writer#/notebooks/40708948/notes/59017778/preview

package question3

func LengthOfLongestSubstring(s string) int {
	freq := [256]int{}

	// 滑动窗口[l...r]
	l, r := 0, -1
	res := 0

	for l != len(s) {
		if r+1 < len(s) && freq[s[r+1]] == 0 {
			r++
			freq[s[r]]++
		} else {
			freq[s[l]]--
			l++
		}

		if res < r-l+1 {
			res = r - l + 1
		}
	}

	return res
}
