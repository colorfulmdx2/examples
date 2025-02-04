package patterns

import (
	"fmt"
	"testing"
)

func TestMonad(t *testing.T) {
	val := Some(10).Map(func(x int) int { return x * 2 }).GetOrElse(0)
	fmt.Println(val)

	noneVal := None[int]().Map(func(x int) int { return x * 2 }).GetOrElse(0)
	fmt.Println(noneVal)
}
