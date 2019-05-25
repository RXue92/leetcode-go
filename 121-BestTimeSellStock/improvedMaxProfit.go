/*
Runtime: 4 ms, faster than 99.21% of Go online submissions for Best Time to Buy and Sell Stock.
Memory Usage: 3.1 MB, less than 60.42% of Go online submissions for Best Time to Buy and Sell Stock.
One pass solution according to reference.
TODO:https://leetcode.com/problems/best-time-to-buy-and-sell-stock/discuss/39038/Kadane's-Algorithm-Since-no-one-has-mentioned-about-this-so-far-%3A)-(In-case-if-interviewer-twists-the-input)
*/

package problem121

func betterSolution(prices []int) int {
	if len(prices) <=1 {
		return 0
	}

	maxUint := ^uint(0)
	maxInt := int(maxUint >> 1)
	minPrice := maxInt
	profit := 0

	for i:=0; i<len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}

		if profit < prices[i] - minPrice {
			profit = prices[i] - minPrice
		}
	}
	return profit
}