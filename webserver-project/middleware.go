package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func CheckAuth() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Checking auth...")
			flag := true
			if flag {
				f(w, r)
				fmt.Println("Auth passed")
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Println("Auth failed")
			}
		}
	}
}

func Logging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()
			fmt.Println("Logging...")
			f(w, r)
			fmt.Println("Logging done")
		}
	}
}
