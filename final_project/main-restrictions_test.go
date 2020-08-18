package main //nolint:golint,stylecheck

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"testing"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
)

func initSevice() {
	//Получить путь к файлу конфигурации
	var cfgPath string
	cfgPath = "BruteforceConfig.yaml"
	//Считать конфигурацию
	yamlFile, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(yamlFile, &Cfg)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	//Задать путь к файлу логирования
	f, err := os.OpenFile(Cfg.LogPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.Println("Запуск http сервера")
	//Запуск сервиса
	handler := &MyHandler{}
	handler.ServiceStarted = false
	Start(handler)
	//Инициализация чёрного и белого списков.
	InitLists()
	InitLrus()
}
func TestTstAllItems(t *testing.T) {
	initSevice()
	t.Run("test login", func(t *testing.T) {
		max := LoginLru.GetLimit()
		for i := 0; i < max; i++ {
			result := TstAllItems("1", "1", "123.123.123.45")
			require.Equal(t, true, result, "Незаконный отказ на шаге "+strconv.Itoa(i))
		}
		result := TstAllItems("1", "1", "123.123.123.45")
		require.Equal(t, false, result, "Незаконное одобрение")
	})
	initSevice()
	t.Run("test password", func(t *testing.T) {
		max := PasswLru.GetLimit()
		for i := 0; i < max; i++ {
			result := TstAllItems("lg1"+strconv.Itoa(i), "1", "123.123.123.45")
			require.Equal(t, true, result, "Незаконный отказ на шаге "+strconv.Itoa(i))
		}
		result := TstAllItems("lg1"+strconv.Itoa(max), "1", "123.123.123.45")
		require.Equal(t, false, result, "Незаконное одобрение")

	})
	initSevice()
	t.Run("test ip", func(t *testing.T) {
		max := IpLru.GetLimit()
		for i := 0; i < max; i++ {
			result := TstAllItems("lg2"+strconv.Itoa(i), strconv.Itoa(i), "123.123.123.45")
			require.Equal(t, true, result, "Незаконный отказ на шаге "+strconv.Itoa(i))
		}

		result := TstAllItems("lg2"+strconv.Itoa(max), strconv.Itoa(max), "123.123.123.45")
		require.Equal(t, false, result, "Незаконное одобрение")

	})
	initSevice()
	t.Run("test drop ip", func(t *testing.T) {
		max := IpLru.GetLimit()
		for i := 0; i < max; i++ {
			result := TstAllItems("lg2"+strconv.Itoa(i), strconv.Itoa(i), "123.123.123.45")
			require.Equal(t, true, result, "Незаконный отказ на шаге "+strconv.Itoa(i))
		}
		DropLogIp("lg21", "123.123.123.45")
		result := TstAllItems("lg2"+strconv.Itoa(max), strconv.Itoa(max), "123.123.123.45")
		require.Equal(t, true, result, "Незаконный отказ")

	})

	initSevice()
	t.Run("test blackList", func(t *testing.T) {
		SetSubnet("123.123.123.45/24", BlackList)
		result := TstAllItems("log", "pwd", "123.123.123.5")
		require.Equal(t, false, result, "Незаконный одобрение")
		DelSubnet("123.123.123.45/24", BlackList)
		result = TstAllItems("log", "pwd", "123.123.123.5")
		require.Equal(t, true, result, "Незаконный отказ")
	})
	initSevice()
	t.Run("test whiteList", func(t *testing.T) {
		SetSubnet("123.123.123.45/24", WhiteList)

		max := LoginLru.GetLimit()
		for i := 0; i < max; i++ {
			result := TstAllItems("1", "1", "123.123.123.45")
			require.Equal(t, true, result, "Незаконный отказ на шаге "+strconv.Itoa(i))
		}
		result := TstAllItems("2", "1", "123.123.123.45")
		require.Equal(t, true, result, "Незаконный отказ")

		DelSubnet("123.123.123.45/24", WhiteList)
		for i := 0; i < max; i++ {
			result := TstAllItems("3", "1", "123.123.123.45")
			require.Equal(t, true, result, "Незаконный отказ на шаге "+strconv.Itoa(i))
		}
		result = TstAllItems("3", "1", "123.123.123.45")
		require.Equal(t, false, result, "Незаконное одобрение")
	})
}
