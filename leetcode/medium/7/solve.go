package _7

const (
    MAX = (1 << 31) - 1
    MIN = -(1 << 31)
)

func reverse(x int) (res int) {

    for x != 0 {
        res = res*10 + x%10
        if res > MAX || res < MIN {
            return 0
        }
        x /= 10
    }
    return
}
