package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

// 上下文是一种类型，她有四种方法，其中三个是与消除和传播有关的方法，与一种值和值传播有关的方案

func theoryContext() {
	// std()
	// out()
	// val()
}

func out() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	sleepAndtalk(ctx, 5*time.Second, "hello")
}

func std() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		fmt.Println("step1")
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		fmt.Println("step2")
		cancel()
		fmt.Println("step3")
	}()
	// time.AfterFunc(time.Second, cancel)
	sleepAndtalk(ctx, 5*time.Second, "hello")
}

var reqID = 123

func val(ctx context.Context, msg string) {
	id, ok := ctx.Value(reqID).(int64)
	if !ok {
		log.Fatal("cloud not find requset ID")
		return
	}
	log.Printf("[%d] %s", id, msg)
}

func decorate(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := rand.Int63()
		ctx = context.WithValue(ctx, reqID, id)
		f(w, r.WithContext(ctx))
	}
}

func sleepAndtalk(ctx context.Context, t time.Duration, s string) {
	select {
	case <-time.After(t):
		fmt.Println(s)
	case <-ctx.Done():
		fmt.Println("done")
		fmt.Println(ctx.Err())
	}
}
