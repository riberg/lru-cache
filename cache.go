package main

import (
	"container/list"
	"sync"
)

type LRUCache interface {
	Add(key, value string) bool
	Get(key string) (value string, ok bool)
	Remove(key string) (ok bool)
}

type entry struct {
	key   string
	value string
}

type LRUCacheImpl struct {
	capacity int
	cache    map[string]*list.Element
	list     *list.List
	mu       sync.Mutex
}

func (l *LRUCacheImpl) Add(key, value string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	if _, ok := l.cache[key]; ok {
		return false
	}

	newEntry := &entry{key: key, value: value}
	element := l.list.PushFront(newEntry)
	l.cache[key] = element

	if l.list.Len() > l.capacity {
		lastElement := l.list.Back()
		if lastElement != nil {
			l.list.Remove(lastElement)
			delete(l.cache, lastElement.Value.(*entry).key)
		}
	}
	return true
}

func (l *LRUCacheImpl) Get(key string) (string, bool) {
	l.mu.Lock()
	defer l.mu.Unlock()

	element, ok := l.cache[key]
	if !ok {
		return "", false
	}

	l.list.MoveToFront(element)
	return element.Value.(*entry).value, true
}

func (l *LRUCacheImpl) Remove(key string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	element, ok := l.cache[key]
	if !ok {
		return false
	}

	l.list.Remove(element)
	delete(l.cache, key)
	return true
}
