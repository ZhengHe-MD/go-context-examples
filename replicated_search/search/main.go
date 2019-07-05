package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type SearchResult struct {
	Hit string
	Err error
}

func searchReplica(query string, replica string) SearchResult {
	log.Printf("search %s on replica %s", query, replica)
	numSecs := 2 + rand.Int63n(8)
	time.Sleep(time.Duration(numSecs)*time.Second)
	log.Printf("search %s on replica %s finished", query, replica)
	return SearchResult{
		Hit: fmt.Sprintf("%s hit in replica %s", query, replica),
		Err: nil,
	}
}

func Search(query string, replicas []string) SearchResult {
	c := make(chan SearchResult, len(replicas))
	for _, replica := range replicas {
		go func() {
			c <- searchReplica(query, replica)
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

	ret := Search("keyword", replicas)
	log.Printf("got search result:%s", ret.Hit)
	log.Println("finish...")

	// pretend there are other works to be done in main routine.
	time.Sleep(10*time.Second)
}

