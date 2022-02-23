package numberof1bits

import (
	"strconv"
)

func hammingWeight(num uint32) int {
	s := strconv.FormatUint(uint64(num), 2)
	count := 0
	for _, ch := range s {
		if ch == '1' {
			count++
		}
	}
	return count
}

func hammingWeight1(num uint32) int {
	count := 0
	for i := 0; i < 32; i++ {
		if 1<<i&num > 0 {
			count++
		}
	}
	return count
}

func hammingWeight2(num uint32) int {
	count := 0
	n := num

	for n > 0 {
		n = n & (n - 1)
		count++
	}
	return count
}
