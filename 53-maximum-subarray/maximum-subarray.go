func maxSubArray(nums []int) int {
    localMax := nums[0]
    globalMax := nums[0]
    
    for i := 1; i < len(nums); i++ {
        if localMax + nums[i] > nums[i] {
            localMax += nums[i]
        } else {
            localMax = nums[i]
        }
        
        if globalMax < localMax {
            globalMax = localMax
        }
    }
    
    return globalMax
}