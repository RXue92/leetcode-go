/*
	Runtime: 0 ms, faster than 100.00% of Go online submissions for Pow(x, n).
	Memory Usage: 2.1 MB, less than 100.00% of Go online submissions for Pow(x, n).
	https://leetcode.com/submissions/detail/226891662/
 */

package problem50

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}

	if n == 1 {
		return x
	}

	if n == -1 {
		return 1/x
	}

	if n == 2 {
		return x*x
	}

	if n % 2 == 0 {
		return myPow(myPow(x, n/2), 2)
	}

	if n < 0 {
		return myPow(x, n+1)/x
	}


	return myPow(x, n-1)*x
}