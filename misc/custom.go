package misc

import (
	"fmt"
)

func ArraysSlices() {
	var s []int
	fmt.Printf("%d, %d, %T, %5t", len(s), cap(s), s, s == nil)	
}