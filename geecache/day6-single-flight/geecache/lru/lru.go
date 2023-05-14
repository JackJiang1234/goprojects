package lru

import "container/list"

type Cache struct {
	maxBytes  int64
	nbytes    int64
	ll        *list.List
	cache     map[string]*list.Element
	OnEvicted func(key string, value Value)
}

type entry struct {
	key   string
	value Value
}

func (e *entry) len() int64 {
	return int64(len(e.key) + e.value.Len())
}

type Value interface {
	Len() int
}

func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

func (c *Cache) Add(key string, value Value) {
	if el, ok := c.cache[key]; ok {
		kv := el.Value.(*entry)
		c.nbytes -= kv.len()
		kv.value = value
		c.nbytes += kv.len()
		c.ll.MoveToFront(el)
	} else {
		kv := &entry{
			key,
			value,
		}
		el := c.ll.PushFront(kv)
		c.cache[key] = el
		c.nbytes += kv.len()
	}
	if c.maxBytes != 0 && c.nbytes >= c.maxBytes {
		c.RemoveOldest()
	}
}

func (c *Cache) Get(key string) (value Value, ok bool) {
	if el, ok := c.cache[key]; ok {
		kv := el.Value.(*entry)
		c.ll.MoveToFront(el)
		return kv.value, true
	} else {
		return nil, false
	}
}

func (c *Cache) RemoveOldest() {
	oldest := c.ll.Back()
	if oldest != nil {
		c.ll.Remove(oldest)
		kv := oldest.Value.(*entry)
		delete(c.cache, kv.key)
		c.nbytes -= kv.len()
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

func (c *Cache) Len() int {
	return c.ll.Len()
}
