package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

var URLReq string
var TimeSleep int

func main() {

	var urlReqCLI = flag.String("url", "", "url")
	var timeSleepCLI = flag.Int("timesleep", 60, "Time sleep")
	flag.Parse()

	URLReq = *urlReqCLI
	TimeSleep = *timeSleepCLI

	t1 := time.Now().Unix()

	if URLReq == "" {
		panic(`NO URL`)
	}

	fmt.Println(URLReq)

	for {

		t2 := time.Now().Unix()

		t := t1 + int64(TimeSleep)
		if t < t2 {
			t1 = t2
			ChekingUrl(URLReq)
		}

		time.Sleep(time.Duration(time.Second * 1))

	}

}

func ChekingUrl(url string) {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(`err`)
	}

	defer resp.Body.Close()
	return

}
