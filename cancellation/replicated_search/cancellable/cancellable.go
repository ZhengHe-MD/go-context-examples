package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type SearchResult struct {
	Hit string
	Err error
}

func searchReplica(ctx context.Context, query string, replica string) SearchResult {
	log.Printf("search %s on replica %s", query, replica)
	c := make(chan SearchResult)
	go func() {
		numSecs := 2 + rand.Int63n(8)
		time.Sleep(time.Duration(numSecs)*time.Second)
		c <-SearchResult{
			Hit: fmt.Sprintf("%s hit in replica %s", query, replica),
			Err: nil,
		}
		log.Printf("search %s on replica %s finished", query, replica)
	}()

	select {
	case <-ctx.Done():
		log.Printf("search %s on replica %s cancelled", query, replica)
		return SearchResult{Err: ctx.Err()}
	case r := <-c:
		return r
	}
}

func Search(ctx context.Context, query string, replicas []string) SearchResult {
	c := make(chan SearchResult, len(replicas))
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	for _, replica := range replicas {
		go func() {
			c <- searchReplica(ctx, query, replica)
		}()
	}
	return <-c
}

func main() {
	replicas := []string{
		"replica1",
		"replica2",
		"replica3",
		"replica4",
		"replica5",
	}

	log.Println("start...")

	ret := Search(context.Background(), "keyword", replicas)
	log.Printf("got search result:%s", ret.Hit)
	log.Println("finish...")

	// pretend there are other works to be done in main routine.
	time.Sleep(10*time.Second)
}
