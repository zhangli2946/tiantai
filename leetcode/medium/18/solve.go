package _18

import "sort"

func fourSum(nums []int, target int) [][]int {
    sort.Ints(nums)
    var (
        lon = len(nums)
        ret = make([][]int, 0)
    )
    for i := 0; i < lon-3; i += 1 {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        for j := i + 1; j < lon-2; j += 1 {
            if j > i+1 && nums[j] == nums[j-1] {
                continue
            }
            for low, high := j+1, lon-1; high > low; {
                n := nums[i] + nums[j] + nums[high] + nums[low]
                switch {
                case n > target:
                    high -= 1
                case n < target:
                    low += 1
                default:
                    for high > low && nums[low] == nums[low+1] {
                        low += 1
                    }
                    for high > low && nums[high] == nums[high-1] {
                        high -= 1
                    }
                    ret = append(ret, []int{nums[i], nums[j], nums[low], nums[high]})
                    low, high = low+1, high-1
                }
            }
        }
    }
    return ret
}
