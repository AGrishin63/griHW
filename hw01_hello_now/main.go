package main

import (
	"fmt"
	"os"
	"time"

	ntp "github.com/beevik/ntp"
)

func main() {
	fmt.Println("current time:\t", time.Now().Format("2006-01-02T15:04:05.999999-07:00"))

	//t.err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	t, err := ntp.Time("0.beevik-nool.ntp.org")

	if err != nil {
		os.Stderr.WriteString("Error on getting exact time!")
		//fmt.Println("Error on getting exact time!")
		return
	}

	fmt.Println("exact time:\t", t.Format("2006-01-02T15:04:05.999999-07:00"))
	// Place your code here
}
