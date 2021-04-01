package _31

import "sort"

func nextPermutation(nums []int) {
    i, o := check(nums)
    if !o {
        sort.Ints(nums)
        return
    }
    solve(nums, i)
    return
}

func solve(ints []int, bIdx int) {
    var (
        aIdx = bIdx - 1
        a    = ints[aIdx]
        cIdx = bIdx
        c    = ints[bIdx]
    )

    for i, curr := range ints[bIdx:] {
        if curr < c && curr > a {
            c, cIdx = curr, i+bIdx
        }
    }
    ints[cIdx] = 0
    sort.Ints(ints[aIdx:])
    ints[aIdx] = c
    return
}

func check(nums []int) (idx int, bool bool) {
    prev := nums[0]
    for i := range nums {
        if prev < nums[i] {
            idx, bool = i, true
        }
        prev = nums[i]
    }
    return
}
