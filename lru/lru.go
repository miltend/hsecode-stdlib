package lru

import (
	"container/list"
)

type Cache struct {
	capacity int
	elements *list.List
	ht       map[int]*list.Element
}

type pair struct {
	key, value int
}

func New(capacity int) *Cache {

	return &Cache{
		capacity: capacity,
		elements: list.New(),
		ht:       make(map[int]*list.Element, capacity),
	}
}

func (cache *Cache) purge() {
	if element := cache.elements.Back(); element != nil {
		item := cache.elements.Remove(element).(pair)
		delete(cache.ht, item.key)
	}
}

func (cache *Cache) Get(key int) (int, bool) {

	elem, ok := cache.ht[key]
	if !ok {
		return 0, false
	}
	value := elem.Value.(pair).value
	cache.elements.MoveToFront(elem)
	return value, true

}

func (cache *Cache) Put(key int, value int) {

	if elem, ok := cache.ht[key]; ok {
		cache.elements.MoveToFront(elem)
		elem.Value = pair{
			key:   key,
			value: value,
		}
	} else {
		if cache.elements.Len() == cache.capacity {
			cache.purge()
		}
		item := pair{
			key:   key,
			value: value,
		}

		elem := cache.elements.PushFront(item)
		cache.ht[key] = elem
	}

}
