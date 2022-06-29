package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	t1 := time.Now().Unix()

	url := "http://ya.ru"
	fmt.Println(url)

	for {

		t2 := time.Now().Unix()

		t := t1 + 60
		if t < t2 {
			t1 = t2
			ChekingUrl(url)
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
