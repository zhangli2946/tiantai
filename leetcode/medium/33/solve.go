package _33

func search(nums []int, target int) int {
    for i, j := 0, len(nums)-1; i <= j; i, j = i+1, j-1 {
        if nums[i] == target {
            return i
        }
        if nums[j] == target {
            return j
        }
    }
    return -1
}
