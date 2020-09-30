package main

import (
	"context"
	"log"
	"math/rand"
	"net/http"
)

type key int

const reqID = key(11)

func Println(ctx context.Context, msg string) {
	id, ok := ctx.Value(reqID).(int64)
	if !ok {
		log.Fatal("cloud not find requset ID")
		return
	}
	log.Printf("[%d] %s", id, msg)
}

func Decorate(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := rand.Int63()
		ctx = context.WithValue(ctx, reqID, id)
		f(w, r.WithContext(ctx))
	}
}
