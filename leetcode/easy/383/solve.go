package _383

func canConstruct(ransomNote string, magazine string) bool {
    var (
        b = [255]int{0}
    )
    for i := range magazine {
        b[magazine[i]] += 1
    }
    for i := range ransomNote {
        b[ransomNote[i]] -= 1
        if b[ransomNote[i]] < 0 {
            return false
        }
    }
    return true
}
