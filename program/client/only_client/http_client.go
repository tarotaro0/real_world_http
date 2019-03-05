package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	values := url.Values{
		"query": {"hello, world"},
	}

	resp, err := http.Get("http://localhost:18888" + "?" + values.Encode())
	// resp, err := http.Head("http://localhost:18888") // Headerのみ取得
	// resp, err := http.PostForm("http://localhost:18888", values) // POSTメソッド
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	log.Println(string(body))
	log.Println("Status:", resp.Status)
	log.Println("StatusCode:", resp.StatusCode)
	log.Println("Header:", resp.Header)
	log.Println("Content-Length:", resp.Header.Get("Content-Length"))
}
