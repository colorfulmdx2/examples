package cache

type Item struct {
	Value any
	Exp   int64
}

type Cache struct {
	Client Client
}

type Client interface {
	Get(key string) (*Item, bool)
	Del(key string)
	Set(key string, val Item)
}

func NewCache(client Client) *Cache {
	return &Cache{Client: client}
}
