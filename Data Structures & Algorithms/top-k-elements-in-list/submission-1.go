func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int, len(nums))
	for _, v := range nums {
		freq[v]++
	}

	buckets := make([][]int, len(nums)+1)
	for num, count := range freq {
		buckets[count] = append(buckets[count], num)
	}

	res := make([]int, 0, k)
	for c := len(buckets) - 1; c >= 1 && len(res) < k; c-- {
		for _, num := range buckets[c] {
			res = append(res, num)
			if len(res) == k {
				break
			}
		}
	}

	return res
}