package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", Decorate(handler))
	http.ListenAndServe("127.0.0.1:1122", nil)
}

func handler(rsp http.ResponseWriter, req *http.Request) {
	cxt := req.Context()
	cxt = context.WithValue(cxt, int(11), int64(100))

	Println(cxt, "handler started")
	defer Println(cxt, "handler ended")

	fmt.Println(cxt.Value("foo"))

	select {
	case <-cxt.Done():
		Println(cxt, "stop!!")
		Println(cxt, cxt.Err().Error())
		http.Error(rsp, cxt.Err().Error(), http.StatusInternalServerError)
	case <-time.After(5 * time.Second):
		Println(cxt, "hello!!")
		fmt.Fprintln(rsp, "hahah")
	}
}
