package solve

import (
    "testing"
)

func TestMergeSortedArray(t *testing.T) {
    merge(
        []int{1, 2, 3, 0, 0, 0},
        3,
        []int{2, 5, 6},
        3,
    )
}

//leetcode submit region begin(Prohibit modification and deletion)
func merge(nums1 []int, m int, nums2 []int, n int) {
    m, n = m-1, n-1
    for index := m + n + 1; index > -1; {
        if m < 0 {
            nums1[index], index, n = nums2[n], index-1, n-1
            continue
        }
        if n < 0 {
            nums1[index], index, m = nums1[m], index-1, m-1
            continue
        }
        if nums1[m] > nums2[n] {
            nums1[index], index, m = nums1[m], index-1, m-1
        } else {
            nums1[index], index, n = nums2[n], index-1, n-1
        }
    }
    return
}

//leetcode submit region end(Prohibit modification and deletion)
/*
func merge(nums1 []int, m int, nums2 []int, n int) {
    for i1, i2 := 0, 0; i2 < n && i1 < m; i1++ {
        for nums1[i1] > nums2[i2] {
            nums1[i1], nums2[i2] = nums2[i2], nums1[i1]
            for i21, i22 := i2, i2+1; i22 < n && nums2[i21] > nums2[i22]; i21, i22 = i22, i22+1 {
                nums2[i21], nums2[i22] = nums2[i22], nums2[i21]
            }
        }
    }
    copy(nums1[m:], nums2[:n])
    return
}

*/
