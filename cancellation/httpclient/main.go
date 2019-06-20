package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 200*time.Millisecond)

	req, _ := http.NewRequest(http.MethodGet, "http://google.com", nil)
	//req, _ := http.NewRequest(http.MethodGet, "http://www.baidu.com", nil)
	req = req.WithContext(ctx)

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		log.Println("Request failed:", err)
		return
	}
	fmt.Println("Response received, status code:", res.StatusCode)
}
