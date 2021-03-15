package _8

const (
    MAX = (1 << 31) - 1
    thr = MAX / 10
    MIN = -(1 << 31)
)

type Handle func(uint8, int, int) (int, int, int)

func myAtoi(s string) int {
    var (
        handleSign       Handle = handleFlag
        handleValue      Handle = record
        handleOther      Handle = terminate
        handleWhiteSpace Handle = ignore

        continu = 0
        sign    = 1
        ans     = 0
    )

    for i := range s {
        switch {
        default:
            sign, ans, continu = handleOther(s[i], sign, ans)
        case '-' == s[i], '+' == s[i]:
            sign, ans, continu = handleSign(s[i], sign, ans)
            handleWhiteSpace, handleSign = terminate, terminate
        case ' ' == s[i]:
            sign, ans, continu = handleWhiteSpace(s[i], sign, ans)
        case '0' <= s[i] && s[i] <= '9':
            sign, ans, continu = handleValue(s[i], sign, ans)
            handleWhiteSpace, handleSign = terminate, terminate
        }
        switch continu {
        default:
            continue
        case 1:
            return ans
        case -1:
            return ans * sign
        }
    }
    return ans * sign
}

func record(u uint8, i4 int, i2 int) (i int, i3 int, b int) {
    if i2 > thr || (i2 == thr && u > '7') {
        if i4 > 0 {
            return i4, MAX, 1
        }
        return i4, MIN, 1
    }
    i3 = i2*10 + int(u-'0')
    return i4, i3, 0
}

func terminate(_ uint8, i4 int, i2 int) (i int, i3 int, b int) { return i4, i2, -1 }
func ignore(_ uint8, i4 int, i2 int) (i int, i3 int, b int)    { return i4, i2, 0 }
func handleFlag(u uint8, _ int, _ int) (i int, i3 int, b int) {
    i, b = 1, 0
    if u == '-' {
        i = - 1
    }
    return
}
