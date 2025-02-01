package slicepkg

// https://ueokande.github.io/go-slice-tricks/

func Cut(slice []any, from int, to int) []any {
	return append(slice[:from], slice[to:]...)
}

func CutGC(slice []any, from int, to int) []any {
	copy(slice[from:], slice[to:])

	defer func(s []any) { // can work with no defer
		for k, n := len(s)-to+from, len(s); k < n; k++ {
			s[k] = nil // or the zero value of T
		}
	}(slice)

	slice = slice[:len(slice)-to+from]

	return slice
}

func Delete(slice []any, i int) []any {
	return append(slice[:i], slice[i+1:]...)

	//a := copy(slice[i:], slice[i+1:])
	//fmt.Println(a)
	//return slice[:i+a]
}

func DeleteGC(slice []any, i int) []any {
	if i < len(slice)-1 {
		copy(slice[i:], slice[i+1:])
	}
	slice[len(slice)-1] = nil // or the zero value of T
	return slice[:len(slice)-1]
}

// DeleteWithoutPreservingOrder delete element with replacing it to last element of slice
// use case when order is not important for the logic
func DeleteWithoutPreservingOrder(slice []any, i int) []any {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func Insert(slice []any, insertElement any, insertIndex int) []any {
	return append(slice[:insertIndex], append([]any{insertElement}, slice[insertIndex:]...)...)
}

func InsertSlice(slice []any, insertSlice []any, insertIndex int) []any {
	return append(slice[:insertIndex], append(insertSlice, slice[insertIndex:]...)...)
}

func Pop(slice []any) (any, []any) {
	popElement, poppedSlice := slice[len(slice)-1], slice[:len(slice)-1]
	return popElement, poppedSlice
}

func Shift(slice []any) (any, []any) {
	shift, shiftedSlice := slice[0], slice[1:]
	return shift, shiftedSlice
}

func Unshift(slice []any, element any) []any {
	return append([]any{element}, slice...)
}

func Extend(slice []any, size int) []any {
	extendedSlice := append(slice, make([]any, size)...)
	return extendedSlice
}
