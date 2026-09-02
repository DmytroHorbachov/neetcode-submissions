func isAnagram(s string, t string) bool {
	count := make(map[rune]int)

	for _, r := range s {
		count[r]++
		fmt.Println(r)
	}

	for _, r := range t {
		count[r]--
		if count[r] < 0 {
			return false
		}
	}

	for _, v := range count {
		if v != 0 {
			return false
		}
	}

	return true
}
