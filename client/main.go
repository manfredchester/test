package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)

	defer cancel()

	ctx = context.WithValue(ctx, "foo", "bar")
	req, err := http.NewRequest(http.MethodGet, "http://localhost:1122", nil)
	if err != nil {
		log.Fatal(err)
	}

	req = req.WithContext(ctx)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(res)
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}
