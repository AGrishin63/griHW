package main

//"fmt"

// func Stop() string {
// 	if LoginLru == nil {
// 		return "Сервис не запущен"
// 	}
// 	fmt.Println("Запустили останов")
// 	stopChan <- true
// 	return "Запустили останов"
// }

// func Start() {
// 	//w := &sync.WaitGroup{}
// 	inChan = make(chan Auth)
// 	outChan = make(chan bool)
// 	stopChan = make(chan bool)
// 	ServiceStarted = true
// 	go func() {
// 		fmt.Println("Запустили горутину")
// 		for {
// 			defer close(inChan)
// 			defer close(stopChan)
// 			select {
// 			case auth, ok := <-inChan:
// 				fmt.Println("Пришёл запрос")
// 				if !ok {
// 					fmt.Println("Запрос не принят")
// 					outChan <- false
// 				}
// 				fmt.Println("Запрос принят.l=" + auth.l)
// 				f := TesAllItems(auth.l, auth.p, auth.i)
// 				fmt.Println("Функция вернула ", f)
// 				outChan <- f //TesAllItems(auth.l, auth.p, auth.i)
// 			case e, f := <-stopChan:
// 				fmt.Println("Сработал stop f=", f, e)
// 				ServiceStarted = false
// 				return
// 			}
// 		}
// 	}()
// 	w.Wait()
// }

// func TestAuth(l string, p string, i string) string {
// 	fmt.Println("Получили запрос")
// 	if !ServiceStarted {
// 		Start()
// 	}
// 	auth := Auth{
// 		l,
// 		p,
// 		i,
// 	}
// 	inChan <- auth
// 	fmt.Println("Отправили запрос")
// 	if <-outChan {
// 		return "true"
// 	}
// 	return "false"
// }

// func main() {
// 	InitLists()
// 	SetSubnet("125.44.33.20/25", WhiteList)
// 	SetSubnet("125.44.33.20/24", WhiteList)
// 	SetSubnet("127.44.33.20/25", BlackList)
// 	SetSubnet("128.44.33.20/24", BlackList)
// 	final, ret := TestIpByLists("125.44.33.27")
// 	fmt.Println(final, ret)

// 	//bList := NewCache(blackListLen)S
// 	return
// }
