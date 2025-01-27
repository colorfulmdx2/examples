package slicepkg

import (
	"fmt"
	"testing"
)

func TestCut(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(s)
	s1 := Cut(s, 4, 8)
	fmt.Println(s1)
	fmt.Println(s)
}

func TestCutGC(t *testing.T) {
	s := []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(s)
	s1 := CutGC(s, 4, 8)
	fmt.Println(s1)
	fmt.Println(s)

}
