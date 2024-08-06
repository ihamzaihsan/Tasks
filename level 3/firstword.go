package piscine

func FirstWord(s string) string {
    start := -1
    for i, char := range s {
        if char != ' ' && start == -1 {
            start = i
        } else if char == ' ' && start != -1 {
            return s[start:i] + "\n"
        }
    }
    return "\n"
}