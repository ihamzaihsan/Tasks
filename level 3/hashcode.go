package piscine

func HashCode(dec string) string {
    lenght := len(dec)
    hashedstr := ""
    for i:=0; i < lenght; i++{
        ass := int(dec[i])
        hashvalue := (ass+lenght) % 127
        if hashvalue < 32 {
            hashvalue = hashvalue + 33
        }
        hashedstr= hashedstr + string(rune(hashvalue))
    }
    return hashedstr
}