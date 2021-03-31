package _29

const max = (1 << 31) - 1

func divide(dividend int, divisor int) int {
    udived = shift
    switch {
    case dividend < 0 && divisor > 0:
        q := udived(-dividend, divisor)
        return -q
    case divisor < 0 && dividend > 0:
        q := udived(dividend, -divisor)
        if q >= max {
            return max
        }
        return -q
    case divisor < 0 && dividend < 0:
        q := udived(-dividend, -divisor)
        if q >= max {
            return max
        }
        return q
    default:
        q := udived(dividend, divisor)
        return q
    }
}

var (
    udived func(int, int) int
)

/*
03/31/2021 22:29	Accepted	0 ms	2.5 MB	golang
03/31/2021 22:10	Accepted	2608 ms	2.5 MB	golang
*/

func shift(d, dd int) (q int) {
    if dd > d {
        return 0
    }
    result, _d, _dd := 1, d, dd
    for ; (_dd << 1) <= _d; {
        result <<= 1
        _dd <<= 1
    }
    return result + udived(d-_dd, dd)
}

func sub(d, dd int) (q int) {
    for d >= dd {
        q += 1
        d -= dd
    }
    return
}
