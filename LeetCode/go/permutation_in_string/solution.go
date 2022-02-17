package permutationinstring

func sameCharacterFrequencies(s1Frequencies map[rune]int, s2 string) bool {
	hashTable := make(map[rune]int)
	for _, s := range s2 {
		hashTable[s]++
	}
	for k, v := range s1Frequencies {
		if v != hashTable[k] {
			return false
		}
	}
	return true
}

func checkInclusion(s1 string, s2 string) bool {
	hashTable := make(map[rune]int)
	for _, s := range s1 {
		hashTable[s]++
	}

	n := len(s1)
	maxI := len(s2) - n
	for i := 0; i <= maxI; i++ {
		subS2 := s2[i : i+n]
		if sameCharacterFrequencies(hashTable, subS2) {
			return true
		}
	}
	return false
}

func checkInclusion1(s1 string, s2 string) bool {
	m := len(s1)
	n := len(s2)
	if m > n {
		return false
	}
	var cnt1, cnt2 [26]int
	for i := range s1 {
		cnt1[s1[i]-'a']++
		cnt2[s2[i]-'a']++
	}
	if cnt1 == cnt2 {
		return true
	}

	for i := m; i < n; i++ {
		cnt2[s2[i]-'a']++
		cnt2[s2[i-m]-'a']--
		if cnt1 == cnt2 {
			return true
		}
	}
	return false
}
