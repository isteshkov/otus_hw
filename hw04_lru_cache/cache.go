package hw04_lru_cache //nolint:golint,stylecheck
import (
	"sync"
)

const defaultCapacity = 5

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	sync.Mutex
	capacity int
	queue    List
	items    map[Key]*listItem
}

type cacheItem struct {
	Key   Key
	Value interface{}
}

func NewCache(capacity int) Cache {
	if capacity <= 0 {
		capacity = defaultCapacity
	}

	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*listItem, capacity),
	}
}

func (c *lruCache) Set(k Key, v interface{}) bool {
	c.Lock()
	defer c.Unlock()

	ci := cacheItem{
		Key:   k,
		Value: v,
	}

	// if key already exists in cache
	if item, ok := c.items[k]; ok {
		c.queue.MoveToFront(item)
		c.queue.Front().Value = ci
		c.items[k] = c.queue.Front()
		return ok
	}

	front := c.queue.PushFront(ci)
	c.items[k] = front

	if c.queue.Len() > 1 {
		key := front.Value.(cacheItem).Key
		c.items[key].Prev = front
	}

	if c.queue.Len() > c.capacity {
		delete(c.items, c.queue.Back().Value.(cacheItem).Key)
		c.queue.Remove(c.queue.Back())
	}

	return false
}

func (c *lruCache) Get(k Key) (interface{}, bool) {
	c.Lock()
	defer c.Unlock()

	if item, ok := c.items[k]; ok {
		c.queue.MoveToFront(item)
		return item.Value.(cacheItem).Value, true
	}

	return nil, false
}

func (c *lruCache) Clear() {
	c.items = make(map[Key]*listItem, c.capacity)
	c.queue = NewList()
}
