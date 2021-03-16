package _11

func max(i, j int) int {
    if i >= j {
        return i
    }
    return j
}
func min(i, j int) int {
    if i < j {
        return i
    }
    return j
}

func maxArea(height []int) (area int) {
    for l, r := 0, len(height)-1; r != l; {
        area = max(area, (r-l)*min(height[l], height[r]))
        if height[l] > height[r] {
            r -= 1
        } else {
            l += 1
        }
    }
    return
}
