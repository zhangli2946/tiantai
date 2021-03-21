package _17

var (
    conversion = map[uint8][]string{
        '2': {"a", "b", "c"},
        '3': {"d", "e", "f"},
        '4': {"g", "h", "i"},
        '5': {"j", "k", "l"},
        '6': {"m", "n", "o"},
        '7': {"p", "q", "r", "s"},
        '8': {"t", "u", "v"},
        '9': {"w", "x", "y", "z"},
    }
)

func letterCombinations(digits string) []string {
    var (
        choiseSet = [][]string{}
    )
    for i := range digits {
        choiseSet = append(choiseSet, conversion[digits[i]])
    }
    return build(choiseSet)
}

func build(set [][]string) []string {
    switch len(set) {
    case 0:
        return []string{}
    case 1:
        return append([]string{}, set[0]...)
    default:
        tmp := build(set[1:])
        ret := make([]string, 0, len(tmp)*len(set[0]))
        for i := range set[0] {
            for j := range tmp {
                ret = append(ret, set[0][i]+tmp[j])
            }
        }
        return ret
    }
}
