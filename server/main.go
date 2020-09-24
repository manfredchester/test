package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("127.0.0.1:1122", nil)
}

func handler(rsp http.ResponseWriter, req *http.Request) {
	fmt.Println("handler started")
	defer fmt.Println("handler ended")

	cxt := req.Context()

	select {
	case <-cxt.Done():
		fmt.Println("stop!!")
		fmt.Println(cxt.Err())
		http.Error(rsp, cxt.Err().Error(), http.StatusInternalServerError)
	case <-time.After(5 * time.Second):
		fmt.Println("hello!!")
		fmt.Fprintln(rsp, "hahah")
	}
}
