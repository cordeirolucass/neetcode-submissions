func isAnagram(s string, t string) bool {
    if len(s) == len(t)  { 
        counts := make(map[rune]int)
        countt := make(map[rune]int)
        for _, char := range s {
            counts[char]++
        }
        for _, char := range t {
            countt[char]++
        }
        for key, value := range counts {
            if countt[key] != value {
                return false
            }
        }
        return true
    }
    return false
}
