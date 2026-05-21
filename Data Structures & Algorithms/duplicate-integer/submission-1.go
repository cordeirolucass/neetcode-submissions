func hasDuplicate(nums []int) bool {
    for i := 0; i < len(nums); i++ {
        var comparing = nums[i]
        for j := 0; j < len(nums); j++ {
            if j == i { continue 
            } else {
                var compared = nums[j]
                if comparing == compared {
                    return true
                }
            }
        }
    }
    return false
}
