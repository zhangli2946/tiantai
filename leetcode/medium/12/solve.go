package _12

var (
    settingStr = [4][3]string{
        {"", "", "M"}, {"M", "D", "C"}, {"C", "L", "X"}, {"X", "V", "I"},
    }
    seq = map[int][]byte{
        9: {2, 0},
        8: {1, 2, 2, 2},
        7: {1, 2, 2},
        6: {1, 2},
        5: {1},
        4: {2, 1},
        3: {2, 2, 2},
        2: {2, 2},
        1: {2},
        0: {},
    }
)

func intToRoman3(num int) string {
    var (
        result string
        para   string
        count  int
        idx    = 4
    )
    for num > 0 {
        count, num, idx, para = num%10, num/10, idx-1, ""
        for i := range seq[count] {
            para += settingStr[idx][seq[count][i]]
        }
        result = para + result
    }
    return result
}

var (
    settings = [4][3]byte{
        {' ', ' ', 'M'}, {'M', 'D', 'C'}, {'C', 'L', 'X'}, {'X', 'V', 'I'},
    }
)

func intToRoman(num int) string {
    var (
        result = make([]byte, 0, 32)
        idx    = 0
        count  int
    )
    for i := 1000; num > 0; i, idx = i/10, idx+1 {
        if num < i {
            continue
        }
        for count = 0; num >= i; num -= i {
            count += 1
        }
        if count > 0 {
            for i := range seq[count] {
                result = append(result, settings[idx][seq[count][i]])
            }
        }
    }
    return string(result)
}

var (
    maps = map[int]string{
        3000: "MMM",
        2000: "MM",
        1000: "M",
        900:  "CM",
        800:  "DCCC",
        700:  "DCC",
        600:  "DC",
        500:  "D",
        400:  "CD",
        300:  "CCC",
        200:  "CC",
        100:  "C",
        90:   "XC",
        80:   "LXXX",
        70:   "LXX",
        60:   "LX",
        50:   "L",
        40:   "XL",
        30:   "XXX",
        20:   "XX",
        10:   "X",
        9:    "IX",
        8:    "VIII",
        7:    "VII",
        6:    "VI",
        5:    "V",
        4:    "IV",
        3:    "III",
        2:    "II",
        1:    "I",
        0:    "",
    }
)

func intToRoman2(num int) string {
    var (
        result string
        max    = 0
    )
    for i := 1000; i > 0; i /= 10 {
        max = i * (num / i)
        num -= max
        result += maps[max]
    }
    return result
}
