package complexnumbermultiplication

import (
	"fmt"
	"strconv"
	"strings"
)

func complexNumberMultiply(num1 string, num2 string) string {
	n1 := strings.Split(num1[0:len(num1)-1], "+")
	n2 := strings.Split(num2[0:len(num2)-1], "+")
	n1Array := make([]int64, len(n1))
	for i, n := range n1 {
		v, _ := strconv.ParseInt(n, 10, 64)
		n1Array[i] = v
	}
	n2Array := make([]int64, len(n2))
	for i, n := range n2 {
		v, _ := strconv.ParseInt(n, 10, 64)
		n2Array[i] = v
	}
	resultArray := make([]int64, 2)
	resultArray[0] = n1Array[0]*n2Array[0] - n1Array[1]*n2Array[1]
	resultArray[1] = n1Array[0]*n2Array[1] + n2Array[0]*n1Array[1]
	return fmt.Sprintf("%d+%di", resultArray[0], resultArray[1])
}
