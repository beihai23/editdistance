package editdistance

import (
	"fmt"
	"testing"
)

func Test_EditDistance(t *testing.T) {
	a := "vim-go"
	b := "hello world"

	ed := EditDistanceDP(a, b)
	edls := EditDistance(a, b)
	fmt.Printf("%s --> %s : %d [dynamic plan]\n", a, b, ed)
	fmt.Printf("%d\n", edls)
	if ed != edls {
		t.Error("EditDistance is wrong")
	}
}
