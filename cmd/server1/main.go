package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

func recoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(write http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("recovered from panic:", r)
				debug.PrintStack()
				http.Error(write, "Something went wrong", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(write, request)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(write http.ResponseWriter, request *http.Request) {
		write.Write([]byte("Hello, World!"))
	})
	mux.HandleFunc("/panic", func(write http.ResponseWriter, request *http.Request) {
		panic("panic")
	})
	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", recoverMiddleware(mux)); err != nil {
		log.Fatalf("Could not listen on port 8080 %v", err)
	}
}
