package helper_lru

import (
	"container/list"
	"sync"
)

type Lru[K comparable, V any] struct {
	lock  sync.RWMutex
	items map[K]*list.Element
	list  *list.List
	size  int
}

type item[K comparable, V any] struct {
	k K
	v V
}

func New[K comparable, V any](size int) *Lru[K, V] {
	return &Lru[K, V]{
		size:  size,
		items: make(map[K]*list.Element),
		list:  list.New(),
	}
}

func (lru *Lru[K, V]) Add(k K, v V) bool {
	defer lru.lock.Unlock()
	lru.lock.Lock()
	if _, ok := lru.items[k]; ok {
		return false
	}
	it := item[K, V]{
		k: k, v: v,
	}
	e := lru.list.PushFront(it)
	lru.items[k] = e
	if lru.list.Len() > lru.size {
		b := lru.list.Back()
		k = b.Value.(item[K, V]).k
		lru.list.Remove(b)
		delete(lru.items, k)
	}
	return true
}

func (lru *Lru[K, V]) Set(k K, v V) {
	defer lru.lock.Unlock()
	lru.lock.Lock()
	if it, ok := lru.items[k]; ok {
		lru.list.Remove(it)
	}
	it := item[K, V]{
		k: k, v: v,
	}
	e := lru.list.PushFront(it)
	lru.items[k] = e
	if lru.list.Len() > lru.size {
		b := lru.list.Back()
		k = b.Value.(item[K, V]).k
		lru.list.Remove(b)
		delete(lru.items, k)
	}
}

func (lru *Lru[K, V]) Get(k K) (V, bool) {
	defer lru.lock.RUnlock()
	lru.lock.RLock()
	it, ok := lru.items[k]
	if ok {
		lru.list.MoveToFront(it)
	}
	if ok {
		return it.Value.(item[K, V]).v, ok
	} else {
		return *new(V), false
	}
}

func (lru *Lru[K, V]) Peek(k K) (V, bool) {
	defer lru.lock.RUnlock()
	lru.lock.RLock()
	it, ok := lru.items[k]
	if ok {
		return it.Value.(item[K, V]).v, ok
	} else {
		return *new(V), false
	}
}

func (lru *Lru[K, V]) Range(f func(k, v any) bool) {
	defer lru.lock.RUnlock()
	lru.lock.RLock()
	for k, it := range lru.items {
		ok := f(k, it.Value.(item[K, V]).v)
		if !ok {
			return
		}
	}
}
