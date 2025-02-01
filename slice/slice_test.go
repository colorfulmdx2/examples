package slicepkg

import (
	"fmt"
	"testing"
)

func TestCut(t *testing.T) {
	s := []any{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(s, "init")
	s1 := Cut(s, 4, 8)
	fmt.Println(s1, "cut")
	fmt.Println(s, "init after cut")
}

func TestCutGC(t *testing.T) {
	s := []any{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(s, "init")
	s1 := CutGC(s, 4, 8)
	fmt.Println(s1, "cut")
	fmt.Println(s, "init after cut")
}

func TestDelete(t *testing.T) {
	s := []any{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(s, "init")
	fmt.Println(Delete(s, 5), "deleted")
	fmt.Println(s, "init after delete")
}

func TestDeleteGC(t *testing.T) {
	s := []any{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(s, "init")
	fmt.Println(DeleteGC(s, 5), "delete")
	fmt.Println(s, "init after delete")
}

func TestDeleteWithoutPreservingOrder(t *testing.T) {
	s := []any{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(s, "init")
	fmt.Println(DeleteWithoutPreservingOrder(s, 5), "delete")
	fmt.Println(s, "init after delete")
}

func TestInsert(t *testing.T) {
	s := make([]any, 0, 100)
	s = append(s, []any{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}...)
	fmt.Println(s, "init")
	fmt.Println(Insert(s, 0, 1), "insert")
	fmt.Println(s, "init after insert, 9 is missing coz len of s is not changed, but 9 in the backed arr")
}

func TestInsertSlice(t *testing.T) {
	s := make([]any, 0, 100)
	s = append(s, []any{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}...)
	fmt.Println(s, "init")
	fmt.Println(InsertSlice(s, []any{0, 0}, 1), "insert")
	fmt.Println(s, "init after insert")
}

func TestPop(t *testing.T) {
	s := []any{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	pop, popSlice := Pop(s)
	fmt.Println(pop, "pop")
	fmt.Println(popSlice, "pop slice")
	fmt.Println(s, "init slice")
}

func TestShift(t *testing.T) {
	s := []any{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	shift, shiftSlice := Shift(s)
	fmt.Println(shift, "shift")
	fmt.Println(shiftSlice, "shift slice")
	fmt.Println(s, "init slice")
}

func TestUnshift(t *testing.T) {
	s := []any{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(s, "init")
	fmt.Println(Unshift(s, 0), "unshift")
	fmt.Println(s, "init after unshift")
}

func TestExtend(t *testing.T) {
	a := make([]any, 3, 5) // len=3, cap=5
	a[0], a[1], a[2] = 1, 2, 3

	fmt.Printf("Before: len=%d, cap=%d, a=%v\n", len(a), cap(a), a)

	a = Extend(a, 2) // Extending by 2

	fmt.Printf("After: len=%d, cap=%d, a=%v\n", len(a), cap(a), a)

}
