package piscine

func RepeatAlpha(s string) string {
    result := ""
    for _, char := range s {
        if char >= 'a' && char <= 'z' {
            repeatCount := int(char - 'a' + 1)
            for i := 0; i < repeatCount; i++ {
                result += string(char)
            }
        } else if char >= 'A' && char <= 'Z' {
            repeatCount := int(char - 'A' + 1)
            for i := 0; i < repeatCount; i++ {
                result += string(char)
            }
        } else {
            result += string(char)
        }
    }
    return result
}
