package main

import (
	"fmt"
	"log"
	"mycache/day3/mycache"
	"net/http"
)

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func main() {
	mycache.NewGroup("scores", 2<<10, mycache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Printf("[slowDB] search for key : %s", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s does not exist", key)
		}))

	addr := "localhost:9999"
	peers := mycache.NewHTTPPool(addr)
	log.Println("mycache is running at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}
