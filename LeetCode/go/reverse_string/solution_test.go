package reversestring

import (
	"fmt"
	"testing"
)

func TestReverseString(t *testing.T) {
	s := []byte("hello")
	reverseString(s)
	fmt.Println(string(s))
}
