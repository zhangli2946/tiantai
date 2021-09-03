package _565

// 本题是不会出现 P 型环的
// 多环 不会出现交点
// 已便利过的元素不必再计算

func arrayNesting2(nums []int) int {
    var max, curr, count int
    for i := range nums {
        for curr, count = nums[i], 0; curr != -1; {
            curr, nums[curr], count = nums[curr], -1, count+1
        }
        if count > max {
            max = count
        }
    }
    return max
}
