package _1235

import "sort"

//
func jobScheduling(startTime []int, endTime []int, profit []int) int {
    startTime, endTime, profit = reorder(startTime, endTime, profit)
    for i := 1; i < len(startTime); i += 1 {
        if firstInvalid := sort.Search(i, func(j int) bool { return endTime[j] > startTime[i] }); firstInvalid == 0 {
            profit[i] = transition(profit[i-1], profit[i])
        } else {
            profit[i] = transition(profit[i-1], profit[firstInvalid-1]+profit[i])
        }
    }
    return profit[len(endTime)-1]
}

func find(startTime int, state map[int]int) (max int) {
    for endTime, dpVal := range state {
        if endTime > startTime {
            continue
        }
        if dpVal > max {
            max = dpVal
        }
    }
    return
}

func transition(i, j int) int {
    if i > j {
        return i
    }
    return j
}

//func jobScheduling(startTime []int, endTime []int, profit []int) int {
//    startTime, endTime, profit = reorder(startTime, endTime, profit)
//    r1 := next(1, 0, startTime, endTime, profit)
//    r2 := next(1, endTime[0], startTime, endTime, profit) + profit[0]
//    if r1 > r2 {
//        return r1
//    }
//    return r2
//}
//
type Sorter struct {
    s, e, p []int
}

func (s Sorter) Len() int { return len(s.s) }

func (s Sorter) Less(i, j int) bool { return s.e[i] < s.e[j] }

func (s Sorter) Swap(i, j int) {
    s.s[i], s.s[j] = s.s[j], s.s[i]
    s.e[i], s.e[j] = s.e[j], s.e[i]
    s.p[i], s.p[j] = s.p[j], s.p[i]
}

func reorder(start []int, end []int, profit []int) ([]int, []int, []int) {
    s := Sorter{start, end, profit}
    sort.Sort(&s)
    return s.s, s.e, s.p
}

//func next(idx, end int, startTime []int, endTime []int, profit []int) (priv int) {
//    if idx >= len(startTime) {
//        return 0
//    }
//    priv = next(idx+1, end, startTime, endTime, profit)
//    r2 := 0
//    if end <= startTime[idx] {
//        r2 = next(idx+1, endTime[idx], startTime, endTime, profit) + profit[idx]
//    }
//    if r2 > priv {
//        priv = r2
//    }
//    return
//}
