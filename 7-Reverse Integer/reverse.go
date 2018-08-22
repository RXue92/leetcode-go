//beat 30%
func reverse(x int) int {
    rev := 0
    isNegative := false
    if x < 0 {
        isNegative = true
        x = -x
    }

    for x > 0 {
        rev = rev*10 + x%10
        x = x/10
    }

    if isNegative {
        rev = -rev
    }

    //check if overflow
    if rev > 1<<31 - 1 || rev < -1<<31 {
        return 0
    }

    return rev
}
