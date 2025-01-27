package cache

import (
	"fmt"
	"testing"
	"time"
)

func TestNewCache(t *testing.T) {

	c1 := NewCache(NewCustomClient())
	c1.Client.Set("a", Item{
		Value: "Aga",
		Exp:   time.Now().UnixMilli(),
	})
	fmt.Println(c1.Client.Get("a"))
	c1.Client.Del("a")
	fmt.Println(c1.Client.Get("a"))
}
