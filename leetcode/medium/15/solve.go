package _15

import "sort"

func threeSum3(nums []int) [][]int {
    sort.Ints(nums)
    var (
        lon = len(nums)
        ans = map[[3]int]byte{}
        ret = make([][]int, 0)
    )

    for i := range nums {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        for j, k := i+1, lon-1; k > j; {
            n := nums[i] + nums[j] + nums[k]
            switch {
            case n > 0:
                k -= 1
            case n < 0:
                j += 1
            default:
                key := [3]int{nums[i], nums[j], nums[k]}
                if _, found := ans[key]; !found {
                    ans[key] = 1
                    ret = append(ret, []int{nums[i], nums[j], nums[k]})
                }
                j, k = j+1, k-1
            }
        }
    }
    return ret
}

// 40 ms	7.2 MB
func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    var (
        lon = len(nums)
        ret = make([][]int, 0)
        ans = map[[3]int]byte{}
    )

    for i := range nums {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        for j, k := i+1, lon-1; k > j; {
            n := nums[i] + nums[j] + nums[k]
            if n > 0 {
                k -= 1
                continue
            }
            if n < 0 {
                j += 1
                continue
            }
            key := [3]int{nums[i], nums[j], nums[k]}
            if _, found := ans[key]; !found {
                ans[key] = 1
                ret = append(ret, []int{nums[i], nums[j], nums[k]})
            }
            j, k = j+1, k-1
        }
    }
    return ret
}

// 48 ms	7 MB
// convert ans  for range big key map has cost!
func threeSum2(nums []int) [][]int {
    sort.Ints(nums)
    var (
        lon = len(nums)
        ans = map[[3]int]byte{}
    )

    for i := range nums {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        for j, k := i+1, lon-1; k > j; {
            n := nums[i] + nums[j] + nums[k]
            if n > 0 {
                k -= 1
                continue
            }
            if n < 0 {
                j += 1
                continue
            }
            ans[[3]int{nums[i], nums[j], nums[k]}] = 1
            j, k = j+1, k-1
        }
    }
    ret := make([][]int, 0, len(ans))
    for k := range ans {
        ret = append(ret, []int{k[0], k[1], k[2]})
    }
    return ret
}
