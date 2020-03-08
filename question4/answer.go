// 4. 寻找两个有序数组的中位数
// https://leetcode-cn.com/problems/median-of-two-sorted-arrays/

package question4

/*
a1 a2 ... ai / ai+1 ai+2 ... a2i
b1 b2 ... bj / bj+1 bj+2 ... b2i

median = (m+n+1)/2
*/

// 思路：使用二分
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)

	// 保证 m <= n
	if m > n {
		nums1, nums2 = nums2, nums1
		m, n = n, m
	}

	var iMin, iMax, halfLen = 0, m, (m + n + 1) / 2
	for {
		var i = (iMin + iMax) / 2
		var j = halfLen - i
		if i < iMax && nums2[j-1] > nums1[i] {
			iMin = i + 1
		} else if i > iMin && nums1[i-1] > nums2[j] {
			iMax = i - 1
		} else { // 匹配到
			maxLeft := 0
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				if nums1[i-1] > nums2[j-1] {
					maxLeft = nums1[i-1]
				} else {
					maxLeft = nums2[j-1]
				}
			}

			if (m+n)%2 == 1 { // 奇数长度
				return float64(maxLeft)
			}

			minRight := 0
			if i == m {
				minRight = nums2[j]
			} else if j == n {
				minRight = nums1[i]
			} else {
				if nums1[i] < nums2[j] {
					minRight = nums1[i]
				} else {
					minRight = nums2[j]
				}
			}

			return float64(maxLeft+minRight) / 2
		}

	} // for
}
