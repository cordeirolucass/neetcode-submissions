func twoSum(nums []int, target int) []int {
    for i:=0; i<len(nums);i++ {
        for j:=0;j < len(nums); j++ {
            if i == j { 
                continue
            } else {
                value := nums[i] + nums[j]
                if value == target {
                    if i < j {
                    answer := []int{i, j}
                    return answer 
                    } else {
                        answer := []int{j, i}
                        return answer
                    } 
                }
            }
        }
    }
    return []int{}
}
