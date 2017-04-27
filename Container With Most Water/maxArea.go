//beat 10%
func maxArea(height []int) int {
    var maxarea, area int
    for i := 0; i < len(height); i++ {
        for j := i + 1; j < len(height); j++ {
            if height[i] > height[j] {
                area = (j - i)*height[j]
            } else {
                area = (j - i)*height[i]
            }
            if area > maxarea {
                maxarea =  area
            }
        }
    }
    return maxarea
}
