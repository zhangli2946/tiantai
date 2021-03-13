package _5

func longestPalindrome(s string) string {
    var (
        m, n = 0, 0
    )
    for i, los := 0, len(s); i < los; i += 1 {
        if h, t := find(s, los, i, i); t-h > n-m {
            m, n = h, t
        }
        if h, t := find(s, los, i, i+1); t-h > n-m {
            m, n = h, t
        }
    }
    return s[m:n]
}

func find(s string, los int, h, t int) (m, n int) {
    for ; h >= 0 && t < los; {
        if s[h] != s[t] {
            return
        }
        m, n = h, t+1
        h, t = h-1, t+1
    }
    return
}
