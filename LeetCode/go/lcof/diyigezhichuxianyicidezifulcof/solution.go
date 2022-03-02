package diyigezhichuxianyicidezifulcof

func firstUniqChar(s string) byte {
	hasTable := make(map[rune]int)
	for _, ch := range s {
		hasTable[ch]++
	}
	for _, ch := range s {
		if hasTable[ch] == 1 {
			return byte(ch)
		}
	}

	return byte(' ')
}
