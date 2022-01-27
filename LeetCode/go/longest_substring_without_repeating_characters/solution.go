package longestsubstringwithoutrepeatingcharacters

func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	start := 0
	end := 0
	visited := map[string]int{}

	for {
		if end >= len(s) {
			currentLength := end - start
			if currentLength > maxLength {
				maxLength = currentLength
			}
			break
		}
		char := string(s[end])
		index, ok := visited[char]
		if !ok {
			visited[char] = end
			end++
		} else {
			currentLength := end - start
			if currentLength > maxLength {
				maxLength = currentLength
			}
			for i := start; i <= index; i += 1 {
				delete(visited, string(s[i]))
			}
			start = index + 1
		}

	}

	return maxLength
}
