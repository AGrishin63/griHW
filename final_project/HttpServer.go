package main

import (
	"fmt"

	"gopkg.in/yaml.v2"

	"io/ioutil"
	"log"
	"net/http"

	n "net/http"
	"net/url"
	"os"
	"sync"
	"time"
)

type MyHandler struct {
	ServiceStarted bool
}

var m *sync.Mutex

func mapRequest(path string, args url.Values, h *MyHandler) string {
	if path == "/testauth" {
		l := args.Get("login")
		p := args.Get("password")
		i := args.Get("ip")
		auth := Auth{
			l,
			p,
			i,
		}
		inChan <- auth
		result := "false"
		if <-outChan {
			result = "true"
		}

		return result
	}
	if path == "/dropbucket" {
		l := args.Get("login")
		i := args.Get("ip")
		m.Lock()
		DropAuthItem(l, LoginLru)
		DropAuthItem(i, IPLru)
		m.Unlock()

		return "true"
	}

	if path == "/blset" {
		sn := args.Get("subnet")
		m.Lock()
		if SetSubnet(sn, BlackList) {
			m.Unlock()

			return "true"
		}
		m.Unlock()

		return "false"
	}

	if path == "/wlset" {
		sn := args.Get("subnet")
		m.Lock()
		if SetSubnet(sn, WhiteList) {
			m.Unlock()

			return "true"
		}
		m.Unlock()

		return "false"
	}

	if path == "/bldrop" {
		sn := args.Get("subnet")
		m.Lock()
		DelSubnet(sn, BlackList)
		m.Unlock()

		return "true"
	}

	if path == "/wldrop" {
		sn := args.Get("subnet")
		m.Lock()
		DelSubnet(sn, WhiteList)
		m.Unlock()

		return "true"
	}

	if path == "/stop" {
		m.Lock()
		s := Stop(h)
		m.Unlock()

		return s
	}

	return "Unknown function call."
}

func (h *MyHandler) ServeHTTP(w n.ResponseWriter, r *n.Request) {
	log.Println("Поступил запрос:", r.URL)
	args := r.URL.Query()
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, mapRequest(r.URL.Path, args, h))
}

// 	if r.URL.Path == "/testauth" {
// 		l := args.Get("login")
// 		p := args.Get("password")
// 		i := args.Get("ip")
// 		w.WriteHeader(http.StatusOK)
// 		auth := Auth{
// 			l,
// 			p,
// 			i,
// 		}
// 		log.Println("Вызов TesAllItems с параметрами:", auth)
// 		m.Lock()
// 		inChan <- auth
// 		result := "false"
// 		m.Unlock()
// 		if <-outChan {
// 			result = "true"
// 		}
// 		log.Println("Резудьтат проверки :", result)
// 		fmt.Fprint(w, result)

// 		return
// 	}
// 	if r.URL.Path == "/dropbucket" {
// 		l := args.Get("login")
// 		i := args.Get("ip")
// 		m.Lock()
// 		DropAuthItem(l, LoginLru)
// 		DropAuthItem(i, IPLru)
// 		m.Unlock()

// 		return
// 	}

// 	if r.URL.Path == "/blset" {
// 		sn := args.Get("subnet")
// 		m.Lock()
// 		if SetSubnet(sn, BlackList) {
// 			fmt.Fprint(w, "true")
// 			m.Unlock()

// 			return
// 		}
// 		fmt.Fprint(w, "false")
// 		m.Unlock()

// 		return
// 	}

// 	if r.URL.Path == "/wlset" {
// 		sn := args.Get("subnet")
// 		m.Lock()
// 		if SetSubnet(sn, WhiteList) {
// 			fmt.Fprint(w, "true")
// 			m.Unlock()

// 			return
// 		}
// 		fmt.Fprint(w, "false")
// 		m.Unlock()

// 		return
// 	}

// 	if r.URL.Path == "/bldrop" {
// 		sn := args.Get("subnet")
// 		m.Lock()
// 		DelSubnet(sn, BlackList)
// 		m.Unlock()
// 		fmt.Fprint(w, "true")

// 		return
// 	}

// 	if r.URL.Path == "/wldrop" {
// 		sn := args.Get("subnet")
// 		m.Lock()
// 		DelSubnet(sn, WhiteList)
// 		m.Unlock()
// 		fmt.Fprint(w, "true")

// 		return
// 	}

// 	if r.URL.Path == "/stop" {
// 		w.WriteHeader(http.StatusOK)
// 		m.Lock()
// 		fmt.Fprint(w, Stop(h))
// 		m.Unlock()

// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprint(w, "Unknown function call.")
// }

func main() {
	m = &sync.Mutex{}

	var cfgPath string
	if len(os.Args) > 1 {
		cfgPath = os.Args[1]
	} else {
		cfgPath = "BruteforceConfig.yaml"
	}
	log.Println("Service cfgPath=", cfgPath)

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
	m.Lock()
	Start(handler)
	m.Unlock()
	//Инициализация чёрного и белого списков.
	InitLists()
	//Таймаут в секундах
	var to time.Duration = 10
	server := &http.Server{
		Addr:         Cfg.Port, //Cfg.port, ":8080"
		Handler:      handler,
		ReadTimeout:  to * time.Second,
		WriteTimeout: to * time.Second,
	}
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

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
				if !ok {
					fmt.Println("Запрос не принят")
					outChan <- false
				}
				outChan <- TstAllItems(auth.l, auth.p, auth.i)
			case e, f := <-stopChan:
				fmt.Println("Сработал stop f=", f, e)
				h.ServiceStarted = false

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
