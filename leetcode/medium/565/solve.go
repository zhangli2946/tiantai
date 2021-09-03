package _565

func arrayNesting(nums []int) int {
    var (
        stack, dp              = []int{0}, make([]int, len(nums))
        curr, depth, max, flag int
    )
    for i := range nums {
        stack, stack[0], depth, flag = stack[:1], nums[i], 0, 0
        for top, direct := 0, 1; top < len(stack) && top > -1; top += direct {
            curr = stack[top] // peak
            switch {
            case direct == -1 && depth != 0:
                dp[stack[top]] = depth
                if stack[top] == flag {
                    depth = 0
                }
            case direct == -1:
                dp[curr] = dp[stack[top+1]] + 1
            case dp[curr] > 0:
                direct = -1
            case instack(stack, nums[curr], &depth):
                dp[stack[top]], flag, direct = depth, nums[curr], -1
            default:
                stack = append(stack, nums[curr])
            }
        }
        if l := dp[nums[i]]; l > max {
            max = l
        }
    }
    return max
}

func instack(stack []int, val int, depth *int) bool {
    if *depth > 0 {
        return true
    }
    for idx, v := range stack {
        if v == val {
            *depth = len(stack) - idx
            return true
        }
    }
    return false
}
