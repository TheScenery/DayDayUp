package findcenterofstargraph

func findCenter(edges [][]int) int {
	c1 := edges[0][0]
	c2 := edges[0][1]
	for _, n := range edges[1] {
		if n == c1 || n == c2 {
			return n
		}
	}
	return 0
}
