package pushdominoes

func calcStatus(dominoes []rune, index int) rune {
	if dominoes[index] != '.' {
		return dominoes[index]
	}
	var lStatus, rStatus rune
	if index-1 >= 0 {
		lStatus = dominoes[index-1]
	}
	if index+1 < len(dominoes) {
		rStatus = dominoes[index+1]
	}
	if lStatus == 'R' && rStatus != 'L' {
		return 'R'
	}
	if rStatus == 'L' && lStatus != 'R' {
		return 'L'
	}
	return dominoes[index]
}

func pushDominoes(dominoes string) string {
	statusChanged := false
	current := []rune(dominoes)
	for {
		next := make([]rune, len(dominoes))
		statusChanged = false
		for i := 0; i < len(dominoes); i++ {
			nextIStatus := calcStatus(current, i)
			if nextIStatus != current[i] {
				statusChanged = true
			}
			next[i] = nextIStatus
		}
		if !statusChanged {
			break
		}
		current = next
	}
	return string(current)
}
