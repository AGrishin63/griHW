package main

import (
	"time"
)

func InitLrus() {
	LoginLru = NewCache(MaxBucketsInCache, Cfg.N)
	PasswLru = NewCache(MaxBucketsInCache, Cfg.M)
	IPLru = NewCache(MaxBucketsInCache, Cfg.K)
}

func TstAllItems(log string, passw string, ip string) bool {
	final, ret := TestIPByLists(ip)
	if final {
		return ret
	}
	if LoginLru == nil {
		InitLrus()
	}

	return testItem(log, LoginLru) && testItem(passw, PasswLru) && testItem(ip, IPLru)
}

func testItem(authItem string, lru Cache) bool {
	now := time.Now()
	tsList, ok := lru.Get(authItem)
	if ok {
		if len(tsList) < lru.GetLimit() { // Событий меньше предела
			lru.Set(authItem, append(tsList, now))

			return true
		}
		// Событий больше предела
		tsList = cleanTSList(tsList, now) // удалить события старше минуты
		lru.Set(authItem, tsList)
		lru.Set(authItem, append(tsList, now))

		return len(tsList) < lru.GetLimit()
	} // Новое событие
	zip := 2 // запас
	tsList = make([]time.Time, 0, lru.GetLimit()+zip)
	tsList = append(tsList, now)
	lru.Set(authItem, tsList)

	return true
}

func cleanTSList(tsList []time.Time, now time.Time) []time.Time {
	tmp := make([]time.Time, 0, cap(tsList))
	for i := len(tsList) - 2; i >= 0; i-- { // Убрать из списка события старше минуты
		if now.Sub(tsList[i]) >= time.Minute {
			tsList = tsList[i+1:]

			break
		}
	}
	if len(tsList) == cap(tsList) { // Если список заполнен, удалить самое старшее событие
		tsList = tsList[1:]
	}
	for i := 0; i < len(tsList); i++ { // переписать для сохранения ёмкости слайса
		tmp = append(tmp, tsList[i])
	}

	return tmp
}

func DropAuthItem(authItem string, lru Cache) {
	lru.DeleteCacheItem(authItem)
}
