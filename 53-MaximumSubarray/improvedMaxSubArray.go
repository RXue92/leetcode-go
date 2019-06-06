/*
DP solution according to https://leetcode.com/problems/maximum-subarray/discuss/20193/DP-solution-and-some-thoughts

Runtime: 8 ms, faster than 26.51% of Go online submissions for Maximum Subarray.
Memory Usage: 3.8 MB, less than 9.49% of Go online submissions for Maximum Subarray.
https://leetcode.com/submissions/detail/234075537/
*/
package problem53

func betterSolution(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	maxDP := dp[0]

	for i := 1; i < len(nums); i++ {
		// dp[i] = nums[i] + dp[i-1]>0?dp[i-1]:0
		dp[i] = nums[i]
		if dp[i-1] > 0 {
			dp[i] += dp[i-1]
		}
		if dp[i] > maxDP {
			maxDP = dp[i] // dp[i] does not equal to maxSubArray(nums[:i])
		}
	}

	return maxDP
}
