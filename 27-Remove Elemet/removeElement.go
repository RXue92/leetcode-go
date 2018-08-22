func removeElement(nums []int, val int) int {
    count := 0
    loc := 0
    k := len(nums)

    for loc + count < k {
        temp := nums[loc]
        if temp == val {
            count++
            nums[loc] = nums[k - count]
            continue
        }
        loc++
    }
    nums = nums[:k-count]
    return k - count

}
