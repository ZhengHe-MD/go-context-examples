package main

import (
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("start processing...")

	ctx := r.Context()

	go func() {
		i := 0
		for i < 10 {
			log.Println(i)
			time.Sleep(1*time.Second)
			i += 1
		}
	}()

	select {
	case <-time.After(2*time.Second):
		_, _ = w.Write([]byte("request processed"))
		log.Println("request processed")
	case <-ctx.Done():
		log.Println("request cancelled")
	}
}

func main() {
	if err := http.ListenAndServe(":8080", http.HandlerFunc(handler)); err != nil {
		log.Panicln(err)
	}
}
