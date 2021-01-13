package hw04_lru_cache //nolint:golint,stylecheck
import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	sync.RWMutex
	capacity int
	queue    List
	items    map[Key]listItem
}

type cacheItem struct {
	Key   Key
	Value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]listItem, capacity),
	}
}

func (c *lruCache) Set(k Key, v interface{}) bool {
	c.Lock()
	defer c.Unlock()

	ci := cacheItem{
		Key:   k,
		Value: v,
	}

	if item, ok := c.items[k]; ok {
		c.queue.MoveToFront(&item)
		c.queue.Front().Value = ci
		c.items[k] = *(c.queue.Front())
		return ok
	}

	c.queue.PushFront(ci)
	c.items[k] = *(c.queue.Front())

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
		c.queue.MoveToFront(&item)
		return item.Value.(cacheItem).Value, true
	}

	return nil, false
}

func (c *lruCache) Clear() {
	c.items = make(map[Key]listItem, c.capacity)
	c.queue = NewList()
}
