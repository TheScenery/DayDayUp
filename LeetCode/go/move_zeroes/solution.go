package movezeros

func moveTo(nums []int, source, target int) {
	for i := source; i < target; i++ {
		nums[i], nums[i+1] = nums[i+1], nums[i]
	}
}

func moveZeroes(nums []int) {
	l := len(nums)
	moveToPosition := l - 1
	for i := l - 1; i >= 0; i-- {
		if nums[i] == 0 {
			if i != moveToPosition {
				moveTo(nums, i, moveToPosition)
			}
			moveToPosition--
		}
	}
}

func moveZeroes1(nums []int) {
	n := len(nums)
	l, r := 0, 0
	for {
		if r >= n {
			return
		}
		if nums[r] != 0 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
		r++
	}
}

func moveZeroes2(nums []int) {
	n := len(nums)
	l, r := 0, 0
	for i := 0; i < n-1; i++ {
		if nums[i] == 0 {
			l = i
			r = i + 1
			break
		}
	}
	for {
		if r >= n {
			return
		}
		if nums[r] != 0 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
		r++
	}
}
