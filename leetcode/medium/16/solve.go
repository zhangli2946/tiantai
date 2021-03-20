package _16

import (
    "sort"
)

func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)
    var (
        lon          = len(nums)-1
        cdiff, pdiff = 0, 1<<63 - 1
        curr, prev   int
    )

    for i := range nums {
        for j, k := i+1, lon; k > j; {
            curr = nums[i] + nums[j] + nums[k]
            cdiff = target - curr
            switch {
            case cdiff > 0:
                j += 1
            case cdiff < 0:
                cdiff = -cdiff
                k -= 1
            default:
                return target
            }
            if pdiff > cdiff {
                pdiff, prev = cdiff, curr
            }
        }
    }
    return prev
}
