func findbig(x int) int {
	// x>10
	num := 10
	for x > num {
		num = num * 10
	}
	return num / 10
}

func isPalindrome(x int) bool {
  //int seems to be int32 in test env
	if x > 1<<31 - 1 {
	    return false
	}

	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	big := findbig(x)

	for x > 0 {
		if x/big != x%10 {
			return false
		}

		x = (x % big) / 10
		big = big / 100
	}
	return true
}
