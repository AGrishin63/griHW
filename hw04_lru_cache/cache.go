package hw04_lru_cache //nolint:golint,stylecheck

type Key string

type Cache interface {
	// Place your code here
	Set(key string, value interface{}) bool // Добавить значение в кэш по ключу
	Get(key string) (interface{}, bool)     // Получить значение из кэша по ключу
	Clear()                                 // Очистить кэш

}

type lruCache struct {
	// Place your code here:
	cap int        // - capacity
	q   *List      // - queue
	itm *cacheItem // - items
}

type cacheItem struct {
	// Place your code here
	k map[string]*ListItem
	l map[*ListItem]string
}

func NewCache(capacity int) Cache {
	cItm := &cacheItem{
		k: make(map[string]*ListItem),
		l: make(map[*ListItem]string),
	}
	return &lruCache{
		cap: capacity,
		q:   NewList(),
		itm: cItm,
	}
}

func (csh *lruCache) Set(key string, value interface{}) bool {
	it, ok := csh.itm.k[key]
	if ok {
		it.Value = value
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

func (csh *lruCache) Get(key string) (interface{}, bool) {
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
