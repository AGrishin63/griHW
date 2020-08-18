package main

import (
	"strconv"
	"strings"
)

type subnet struct {
	ipSub   string
	maskLen int
}

func InitBList() {
	BlackList = make(map[string]subnet)
}

func InitWList() {
	WhiteList = make(map[string]subnet)
}

func InitLists() {
	InitBList()
	InitWList()
}

func SetSubnet(sn string, list map[string]subnet) bool {
	sl := strings.Split(sn, "/")
	len, err := strconv.Atoi(sl[1])
	if err != nil {
		return false
	}
	list[sn] = subnet{
		ipSub:   getSubnetIP(sl[0], len),
		maskLen: len}

	return true
}

func DelSubnet(sn string, list map[string]subnet) {
	delete(list, sn)
}

func getSubnetIP(ip string, maskLen int) string {
	sn := ""
	//Кол. чисел в IP адресе
	ipMax := 4
	for i, str := range strings.Split(ip, ".") {
		d4, _ := strconv.Atoi(str)
		sn += strconv.Itoa(d4 & int(masks[maskLen][i]))
		if i != ipMax-1 {
			sn += "."
		}
	}

	return sn
}

func isItSameSubNet(ip string, listSubName string, maskLen int) bool {
	return getSubnetIP(ip, maskLen) == listSubName
}

func TestIPByList(ip string, list map[string]subnet) bool {
	for key := range list {
		if isItSameSubNet(ip, list[key].ipSub, list[key].maskLen) {
			return true
		}
	}

	return false
}
func TestIPByLists(ip string) (bool, bool) {
	if TestIPByList(ip, WhiteList) {
		return true, true
	}
	if TestIPByList(ip, BlackList) {
		return true, false
	}

	return false, false
}
