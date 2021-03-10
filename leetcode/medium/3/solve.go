package _3

func lengthOfLongestSubstring(s string) int {
    return lengthOfLongestSubstring_2(s)
}

func cleanStat(bitMap *[255]int, s string) int {
    for _, c := range s {
        (*bitMap)[c] = -1
    }
    return len(s)
}

// 4ms 2.6M
func lengthOfLongestSubstring_1(s string) int {
    var (
        tail   int
        record = 0
        stat   = [255]int{}
        l      = len(s)
        curr   = 0
    )
    for ; curr < l; curr += 1 {
        char := s[curr]
        prev := stat[char]
        if prev != 0 {
            if n := curr - tail; n > record {
                record = n
            }
            cleanStat(&stat, s[tail:prev])
            tail = prev
        }
        stat[char] = curr + 1
    }
    if n := curr - tail; n > record {
        return n
    }
    return record
}

// 0ms 2.6M
func lengthOfLongestSubstring_2(s string) int {
    var (
        tail   int
        record = 0
        bitMap = [255]int{}
        l      = len(s)
        curr   = 0
    )
    for ; curr < l; curr += 1 {
        char := s[curr]
        prev := bitMap[char]
        if prev == 0 {
            bitMap[char] = curr + 1
            continue
        }
        if n := curr - tail; n > record {
            record = n
        }
        cleanStat(&bitMap, s[tail:prev])
        tail = prev
        bitMap[char] = curr + 1
    }
    if n := curr - tail; n > record {
        return n
    }
    return record
}
