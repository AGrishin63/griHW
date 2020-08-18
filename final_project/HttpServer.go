package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	n "net/http"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type MyHandler struct {
	ServiceStarted bool
}

func (h *MyHandler) ServeHTTP(w n.ResponseWriter, r *n.Request) {
	log.Println("Поступил запрос:", r.URL)
	args := r.URL.Query()
	if r.URL.Path == "/testauth" {
		l := args.Get("login")
		p := args.Get("password")
		i := args.Get("ip")
		w.WriteHeader(http.StatusOK)
		auth := Auth{
			l,
			p,
			i,
		}
		log.Println("Вызов TesAllItems с параметрами:", auth)
		inChan <- auth
		result := "false"
		if <-outChan {
			result = "true"
		}
		log.Println("Резудьтат проверки :", result)
		fmt.Fprint(w, result)
		return
	}
	if r.URL.Path == "/dropbucket" {
		l := args.Get("login")
		i := args.Get("ip")
		DropAuthItem(l, LoginLru)
		DropAuthItem(i, IpLru)
		return
	}

	if r.URL.Path == "/blset" {
		sn := args.Get("subnet")
		if SetSubnet(sn, BlackList) {
			fmt.Fprint(w, "true")
			return
		}
		fmt.Fprint(w, "false")
		return
	}

	if r.URL.Path == "/wlset" {
		sn := args.Get("subnet")
		if SetSubnet(sn, WhiteList) {
			fmt.Fprint(w, "true")
			return
		}
		fmt.Fprint(w, "false")
		return
	}

	if r.URL.Path == "/bldrop" {
		sn := args.Get("subnet")
		DelSubnet(sn, BlackList)
		fmt.Fprint(w, "true")
		return
	}

	if r.URL.Path == "/wldrop" {
		sn := args.Get("subnet")
		DelSubnet(sn, WhiteList)
		fmt.Fprint(w, "true")
		return
	}

	if r.URL.Path == "/stop" {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, Stop(h))
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Unknown function call.")

}

func main() {
	//Получить путь к файлу конфигурации
	var cfgPath string
	if len(os.Args) > 1 {
		cfgPath = os.Args[1]
	} else {
		cfgPath = "BruteforceConfig.yaml"
	}
	log.Println("Service cfgPath=", cfgPath)
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
	//
	server := &http.Server{
		Addr:         Cfg.Port, //Cfg.port, ":8080"
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server.ListenAndServe()
}

func Start(h *MyHandler) {
	inChan = make(chan Auth)
	outChan = make(chan bool)
	stopChan = make(chan bool)
	h.ServiceStarted = true
	go func(h *MyHandler) {
		log.Println("Запустили горутину")
		for {
			select {
			case auth, ok := <-inChan:
				//fmt.Println("Пришёл запрос")
				if !ok {
					fmt.Println("Запрос не принят")
					outChan <- false
				}
				//fmt.Println("Запрос принят.l=" + auth.l)
				//f := TesAllItems(auth.l, auth.p, auth.i)
				//fmt.Println("Функция вернула ", f)
				outChan <- TstAllItems(auth.l, auth.p, auth.i)
			case e, f := <-stopChan:
				fmt.Println("Сработал stop f=", f, e)
				h.ServiceStarted = false
				//fmt.Println("Сервис остановлен")
				return
			}
		}
	}(h)

}

func Stop(h *MyHandler) string {
	if !h.ServiceStarted {
		return "Сервис не запущен"
	}
	log.Println("Запустили останов")
	stopChan <- true
	return "Сервис остановлен"
}
