package slicepkg

func Cut(s []int, i int, j int) []int {
	return append(s[:i], s[j:]...)
}

func CutGC(s []any, i int, j int) []any {
	copy(s[i:], s[j:])

	defer func(s []any) { // can work with no defer
		for k, n := len(s)-j+i, len(s); k < n; k++ {
			s[k] = nil // or the zero value of T
		}
	}(s)

	s = s[:len(s)-j+i]

	return s
}
