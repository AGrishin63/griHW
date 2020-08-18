package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main2() {
	var s string = ""
	var url string = ""

	fmt.Println("Привет. Вы можете ввести сдудующие команды:")
	fmt.Println("addB-<ip/mask>  - добавить подсеть в чёрный список")
	fmt.Println("addW-<ip/mask>  - добавить подсеть в белый список")
	fmt.Println("dropB-<ip/mask> - удалить подсеть из чёрного списка")
	fmt.Println("dropW-<ip/mask> - удалить подсеть из белого списка")
	fmt.Println("exit            - завершить работу")
	fmt.Println("<ip/mask> -  это маска подсети в виде 127.721.12.16.24")

	for {
		fmt.Println("Введите команду:")
		fmt.Fscan(os.Stdin, &s)
		if strings.Contains(strings.ToLower(s), "exit") {
			return
		}

		sl := strings.Split(s, "-")
		if len(sl) != 2 {
			fmt.Println("Ошибка ввода. надо ввести команду и через тире маску подсети.")
		}
		fmt.Println(sl[0], sl[1])
		if strings.Contains(strings.ToLower(sl[0]), "addb") {
			url = "http://localhost:8080/blset?subnet=" + sl[1]
		}
		if strings.Contains(strings.ToLower(sl[0]), "addw") {
			url = "http://localhost:8080/wlset?subnet=" + sl[1]
		}
		if strings.Contains(strings.ToLower(sl[0]), "dropb") {
			url = "http://localhost:8080/bldrop?subnet=" + sl[1]
		}
		if strings.Contains(strings.ToLower(sl[0]), "dropw") {
			url = "http://localhost:8080/wldrop?subnet=" + sl[1]
		}

		fmt.Println(url)
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	}
}
