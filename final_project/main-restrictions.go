package main

import (
	"time"
)

func InitLrus() {
	LoginLru = NewCache(MaxBucketsInCache, Cfg.N)
	PasswLru = NewCache(MaxBucketsInCache, Cfg.M)
	IpLru = NewCache(MaxBucketsInCache, Cfg.K)
}

func TstAllItems(log string, passw string, ip string) bool {
	final, ret := TestIpByLists(ip)
	if final {
		return ret
	}
	if LoginLru == nil {
		InitLrus()
	}
	return testItem(log, LoginLru) && testItem(passw, PasswLru) && testItem(ip, IpLru)
}

func testItem(authItem string, lru Cache) bool {
	now := time.Now()
	tsList, ok := lru.Get(authItem)
	//fmt.Println("len(tsList)=", len(tsList))
	//fmt.Println("lru.GetLimit()=", lru.GetLimit())
	if ok {
		if len(tsList) < lru.GetLimit() { //Событий меньше предела
			lru.Set(authItem, append(tsList, now))
			return true
		} else { // Событий больше предела
			tsList = cleanTsList(tsList, now) // удалить события старше минуты
			lru.Set(authItem, tsList)
			lru.Set(authItem, append(tsList, now))
			if len(tsList) < lru.GetLimit() {
				return true
			} else {
				return false
			}
		}
	} else { //Новое событие

		tsList = make([]time.Time, 0, lru.GetLimit()+2)
		tsList = append(tsList, now)
		lru.Set(authItem, tsList)
		//fmt.Println("Создали tsList c len=", len(tsList))
		return true
	}
}

func cleanTsList(tsList []time.Time, now time.Time) []time.Time {
	tmp := make([]time.Time, 0, cap(tsList))
	for i := len(tsList) - 2; i >= 0; i-- { //Убрать из списка события старше минуты
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
// func DropLogIp(login string, ip string) {
// 	DropAuthItem(login, LoginLru)
// 	DropAuthItem(ip, IpLru)
// }
