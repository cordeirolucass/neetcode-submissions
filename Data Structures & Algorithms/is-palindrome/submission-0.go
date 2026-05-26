func limparString(s string) string {
    var builder strings.Builder

    for _, r:= range s {
        if unicode.IsLetter(r) || unicode.IsDigit(r) {
            builder.WriteRune(unicode.ToLower(r))
        }
    }
    return builder.String()
}
func isPalindrome(s string) bool {
    if len(s) == 0 { return false
    } else { 
        strLimpa := limparString(s)
        n := len(strLimpa)
        for i := 0; i < n/2; i++ {
            if strLimpa[i] != strLimpa[n-i-1] {
                return false
            }
        }
    }
    return true
}
