package q42

import (
    "testing"
)

func TestTrappingRainWater(t *testing.T) {
    _ = trap
}

//leetcode submit region begin(Prohibit modification and deletion)
func trap(height []int) (area int) {
    for l, r, lm, rm := 0, len(height)-1, height[0], height[len(height)-1]; l <= r; {
        if height[l] <= height[r] {
            if lm < height[l] {
                lm = height[l]
            }
            area, l = area+lm-height[l], l+1
            continue
        }
        if rm < height[r] { rm = height[r] }
        area, r = area+rm-height[r], r-1
    }
    return
}

//leetcode submit region end(Prohibit modification and deletion)
