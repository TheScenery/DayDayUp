package pancakesorting

func pancake(arr []int, k int) {
	for i, j := 0, k; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func pancakeSort(arr []int) []int {
	result := make([]int, 0)
	n := len(arr)
	for v := n; v > 0; v-- {
		vIndex := 0
		for i := v - 1; i >= 0; i-- {
			if arr[i] == v {
				vIndex = i
				break
			}
		}
		if vIndex != v-1 {
			pancake(arr, vIndex)
			result = append(result, vIndex+1)
			pancake(arr, v-1)
			result = append(result, v)
		}
	}
	return result
}
