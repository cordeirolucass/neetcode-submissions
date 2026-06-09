func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var resultadoFinal [][]int = nil
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {

			continue
		}
		left := i+1
		right := len(nums)-1
		for left < right {
			var resultado []int = nil
			soma := nums[i] + nums[left] + nums[right]
			if soma > 0 {
				right -= 1
			} else if soma < 0 {
				left +=1
			} else {
				resultado = []int{nums[i], nums[left], nums[right]}
				right -= 1
				left += 1
				for left < right && nums[left] == nums[left-1]{
					left += 1
				}
				for left < right && nums[right] == nums[right+1] {
					right -= 1
				}
				resultadoFinal = append(resultadoFinal, resultado)
			} 
		}
	}
	return resultadoFinal
}
