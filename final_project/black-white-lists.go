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
		ipSub:   getSubnetIp(sl[0], len),
		maskLen: len}
	//fmt.Println("после list[sn]=", list[sn])
	return true
}

func DelSubnet(sn string, list map[string]subnet) {
	delete(list, sn)
}

func getSubnetIp(ip string, maskLen int) string {
	sn := ""
	for i, str := range strings.Split(ip, ".") {
		d4, _ := strconv.Atoi(str)
		//fmt.Println(d4, int(masks[maskLen][i]))
		sn += strconv.Itoa(d4 & int(masks[maskLen][i]))
		if i != 3 {
			sn += "."
		}
	}
	return sn
}

func isItSameSubNet(ip string, listSubName string, maskLen int) bool {
	//fmt.Println(getSubnetIp(ip, maskLen), listSubName)
	return getSubnetIp(ip, maskLen) == listSubName
}

func TestIpByList(ip string, list map[string]subnet) bool {
	for key := range list {
		if isItSameSubNet(ip, list[key].ipSub, list[key].maskLen) {
			return true
		}
	}
	return false
}
func TestIpByLists(ip string) (bool, bool) {
	if TestIpByList(ip, WhiteList) {
		return true, true
	}
	if TestIpByList(ip, BlackList) {
		return true, false
	}
	return false, false
}
