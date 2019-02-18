package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true) // 第２引数がtrueの時は、リクエストのbodyをそのまま出力
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))                               // request dump の表示
	fmt.Fprintf(w, "<html><body>hellooooo</body></html>\n") // responseに書き込み
}

func main() {
	var httpServer http.Server
	http.HandleFunc("/", handler) // path "/" にアクセス時、handler()をcall
	log.Println("start http listening :18888")
	httpServer.Addr = ":18888"               // 18888 portに設定
	log.Println(httpServer.ListenAndServe()) // 設定したportで待ち受け、handlerを呼び出すことを開始
}
