package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	var item *ListItem
	var found bool

	if item, found = c.items[key]; !found {
		if c.capacity <= c.queue.Len() {
			backItem := c.queue.Back()
			delete(c.items, backItem.Key)

			c.queue.Remove(backItem)
		}
		newItem := c.queue.PushFront(value)
		newItem.Key = key
		c.items[key] = newItem
		return false
	}
	item.Value = value
	c.queue.MoveToFront(item)
	return true
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	if item, found := c.items[key]; found {
		c.queue.MoveToFront(item)
		return item.Value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
