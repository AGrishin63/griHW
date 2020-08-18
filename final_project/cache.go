package main 
import (
	"time"
)

//type Key string

type Cache interface {
	// Place your code here
	Set(key string, value []time.Time) bool // Добавить значение в кэш по ключу
	Get(key string) ([]time.Time, bool)     // Получить значение из кэша по ключу
	GetLimit() int                          // Получить значение ограничения
	Clear()                                 // Очистить кэш
	DeleteCacheItem(key string)             // Удалить элемент из кэша
}

type lruCache struct {
	cap   int        // - capacity
	limit int        // max items number
	q     *List      // - queue
	itm   *cacheItem // - items
}

type cacheItem struct {
	// Place your code here
	k map[string]*ListItem
	l map[*ListItem]string
}

func NewCache(capacity int, limit int) Cache {
	cItm := &cacheItem{
		k: make(map[string]*ListItem),
		l: make(map[*ListItem]string),
	}

	return &lruCache{
		cap:   capacity,
		limit: limit,
		q:     NewList(),
		itm:   cItm,
	}
}
func (csh *lruCache) GetLimit() int {
	return csh.limit
}

func (csh *lruCache) Set(key string, value []time.Time) bool {
	it, ok := csh.itm.k[key]
	if ok {
		it.Value = value
		csh.itm.k[key] = it
		// fmt.Println("value=", value)
		// fmt.Println("it=", it)
		// fmt.Println("From cashe=", csh.itm.k[key])
		csh.q.MoveToFront(it)

		return true
	}
	csh.itm.k[key] = csh.q.PushFront(value)
	csh.itm.l[csh.itm.k[key]] = key
	if csh.q.Size > csh.cap {
		delete(csh.itm.k, csh.itm.l[csh.q.Back()])
		delete(csh.itm.l, csh.q.Back())
		csh.q.Remove(csh.q.Back())
	}

	return false
}

func (csh *lruCache) Get(key string) ([]time.Time, bool) {
	it, ok := csh.itm.k[key]
	if ok {
		csh.q.MoveToFront(it)
		return it.Value, true
	}
	
	return nil, false
}

func (csh *lruCache) Clear() {
	for i := 0; i < csh.q.Size; i++ {
		csh.q.Remove(csh.q.First)
	}
	for key := range csh.itm.k {
		delete(csh.itm.k, key)
	}
	for ls := range csh.itm.l {
		delete(csh.itm.l, ls)
	}
}
func (csh *lruCache) DeleteCacheItem(key string) {
	delete(csh.itm.l, csh.itm.k[key])
	csh.q.Remove(csh.itm.k[key])
	delete(csh.itm.k, key)
}
