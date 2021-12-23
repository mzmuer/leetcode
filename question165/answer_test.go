// 165. 比较版本号
// https://leetcode-cn.com/problems/compare-version-numbers/

package question165

import (
	"strconv"
	"strings"
	"testing"
)

func Test_compareVersion(t *testing.T) {
	t.Log(compareVersion("1.0.1", "1"))
}

func compareVersion(version1 string, version2 string) int {
	v1Arr := strings.Split(version1, ".")
	v2Arr := strings.Split(version2, ".")

	for i := 0; i < len(v1Arr) || i < len(v2Arr); i++ {
		v1, v2 := 0, 0

		if i < len(v1Arr) {
			v1, _ = strconv.Atoi(v1Arr[i])
		}
		if i < len(v2Arr) {
			v2, _ = strconv.Atoi(v2Arr[i])
		}

		if v1 > v2 {
			return 1
		} else if v1 < v2 {
			return -1
		}
	}

	return 0
}
