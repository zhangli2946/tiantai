package _34

func searchRange(nums []int, target int) []int {
    head, tail := 0, len(nums)-1
    if tail == -1 {
        return []int{-1, -1}
    }
    for ; tail > head; {
        if mid := (tail + head) / 2; target > nums[mid] {
            head = mid + 1
        } else {
            tail = mid
        }
    }
    if target != nums[head] {
        return []int{-1, -1}
    }
    start := head
    head, tail = 0, len(nums)-1
    for ; tail > head; {
        if mid := (tail+head)/2 + 1; target < nums[mid] {
            tail = mid - 1
        } else {
            head = mid
        }
    }

    if target != nums[tail] {
        return []int{-1, -1}
    }
    return []int{start, tail}
}
