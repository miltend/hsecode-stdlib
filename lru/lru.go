package lru

//package main

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
		ht:       make(map[int]*list.Element),
	}
}

func (cache *Cache) purge() {
	if element := cache.elements.Back(); element != nil {
		item := cache.elements.Remove(element).(*pair)
		delete(cache.ht, item.key)
	}
}

func (cache *Cache) Get(key int) (int, bool) {

	elem, ok := cache.ht[key]
	if !ok {
		return 0, false
	}
	cache.elements.MoveToFront(elem)
	return elem.Value.(*pair).value, true

}

func (cache *Cache) Put(key int, value int) {

	//if elem, ok := cache.ht[key]; ok == true {
	elem, _ := cache.ht[key]
	cache.elements.MoveToFront(elem)
	elem.Value.(*pair).value = value
	//}

	if cache.elements.Len() == cache.capacity {
		cache.purge()
	}

	item := &pair{
		key:   key,
		value: value,
	}

	elem = cache.elements.PushFront(item)
	cache.ht[item.key] = elem
}

//
//func main() {
//
//
//}
