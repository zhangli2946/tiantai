package _1

func twoSum(nums []int, target int) []int {
    var (
        lon         = len(nums) - 1
        i, j        int
    )
    for i = range nums {
        for j = lon; i < j; j -= 1 {
            if target == nums[i] + nums[j] {
                return []int{i, j}
            }
        }
    }
    return []int{i, j}
}
