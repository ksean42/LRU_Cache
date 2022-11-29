package LRUCache

import (
	"sync"
)

type LRUCache struct {
	cap   uint
	len   uint
	data  map[interface{}]interface{}
	queue *List
	sync.RWMutex
}

func NewLRUCache(cap uint) *LRUCache {
	return &LRUCache{
		cap:   cap,
		len:   0,
		queue: NewList(),
		data:  make(map[interface{}]interface{}, cap),
	}
}

func (c *LRUCache) Len() uint {
	c.RLock()
	defer c.RUnlock()
	return c.len
}

func (c *LRUCache) Cap() uint {
	c.RLock()
	defer c.RUnlock()
	return c.cap
}

func (c *LRUCache) Add(key, value interface{}) {
	c.Lock()
	defer c.Unlock()
	_, ok := c.data[key]
	if ok {
		c.queue.Delete(key)
		c.data[key] = value
		c.queue.PushBack(key)
		return
	}
	if c.cap == c.len {
		v := c.queue.PopFront()
		delete(c.data, v)
	} else {
		c.len++
	}
	c.data[key] = value
	c.queue.PushBack(key)
}

func (c *LRUCache) Get(key interface{}) (interface{}, bool) {
	c.RLock()
	defer c.RUnlock()
	if value, ok := c.data[key]; ok {
		c.queue.Delete(key)
		c.queue.PushBack(key)
		return value, true
	}
	return nil, false
}

func (c *LRUCache) Remove(key interface{}) {
	c.Lock()
	defer c.Unlock()
	if _, ok := c.data[key]; ok {
		delete(c.data, key)
		c.queue.Delete(key)
		c.len--
	}
}

func (c *LRUCache) Clear() {
	c.Lock()
	c.len = 0
	c.data = make(map[interface{}]interface{}, c.cap)
	c.queue = NewList()
	c.Unlock()
}

func (c *LRUCache) GetAll() (map[interface{}]interface{}, List) {
	return c.data, *c.queue
}
