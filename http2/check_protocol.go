package main

import (
	"fmt"
	"net/http"
)

// GODEBUG=http2client=0 go run => Disable cleint http2 support
// GODEBUG=http2server=0 go run => Disable server http2 support
// GODEBUG=http2debug=1 go run => Enable debug log related to http2
// GODEBUG=http2debug=2 go run => Output logs more than debug=1
func main() {
	resp, err := http.Get("http://gogole.com/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Printf("Protocol Version: %s\n", resp.Proto)
}
