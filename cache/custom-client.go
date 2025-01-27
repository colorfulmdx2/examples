package cache

import (
	"fmt"
	"sync"
)

type CustomClient struct {
	mu    sync.RWMutex
	store map[string]Item
}

type CacheCustomClient interface {
	Get(key string) (*Item, bool)
	Del(key string)
	Set(key string, val Item)
}

func (c *CustomClient) Get(key string) (*Item, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, exists := c.store[key]

	return &item, exists
}
func (c *CustomClient) Del(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.store, key)
	fmt.Println(key, "is deleted")
}
func (c *CustomClient) Set(key string, val Item) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.store[key] = val
}
func NewCustomClient() *CustomClient {
	return &CustomClient{store: make(map[string]Item)}
}
