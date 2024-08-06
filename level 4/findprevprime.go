package piscine

func FindPrevPrime(nb int) int {
    isPrime := func(n int) bool {
        if n <= 1 {
            return false
        }
        for i := 2; i*i <= n; i++ {
            if n%i == 0 {
                return false
            }
        }
        return true
    }

    for nb >= 2 {
        if isPrime(nb) {
            return nb
        }
        nb--
    }
    return 0
}
