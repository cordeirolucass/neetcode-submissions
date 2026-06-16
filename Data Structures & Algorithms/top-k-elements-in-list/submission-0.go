func topKFrequent(nums []int, k int) []int {
	bucket := make([][]int, len(nums)+1)
	count := make(map[int]int)
	var result []int
	for _, num := range nums {
		count[num]++
	}

	for number, freq := range count {
		bucket[freq] = append(bucket[freq], number)
	}
	for i := len(bucket)-1; len(result) < k; i-- {
		result = append(result, bucket[i]...)
	}
return result
}
