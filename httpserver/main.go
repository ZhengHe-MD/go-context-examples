package main

import (
	"log"
	"net/http"
	"time"
)

func fooHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("start foo...")

	for i := 0; i < 10; i++ {
		log.Println(i)
		time.Sleep(1*time.Second)
	}
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("start bar...")

	Loop:
	for i := 0; i < 10; i++ {
		log.Println(i)
		select {
		case <-time.After(1*time.Second):
			// do nothing
		case <-r.Context().Done():
			break Loop
		}
	}
}

func bazHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("start baz...")

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
	http.HandleFunc("/foo", fooHandler)
	http.HandleFunc("/bar", barHandler)
	http.HandleFunc("/baz", bazHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Panicln(err)
	}
}
