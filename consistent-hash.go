package main

import (
	"flag"
	"fmt"
	"log"
	"stathat.com/c/consistent"
	"strconv"
)

var keysPtr = flag.Int("keys", 10000, "key number")
var nodesPtr = flag.Int("nodes", 3, "node number of old cluster")
var newNodesPtr = flag.Int("new-nodes", 4, "node number of new cluster")

func hash(key int, nodes int) int {
	return key % nodes
}

func main() {

	flag.Parse()
	var keys = *keysPtr
	var nodes = *nodesPtr
	var newNodes = *newNodesPtr

	c := consistent.New()
	for i := 0; i < nodes; i++ {
		c.Add(strconv.Itoa(i))
	}

	newC := consistent.New()
	for i := 0; i < newNodes; i++ {
		newC.Add(strconv.Itoa(i))
	}

	migrate := 0
	for i := 0; i < keys; i++ {
		server, err := c.Get(strconv.Itoa(i))
		if err != nil {
			log.Fatal(err)
		}

		newServer, err := newC.Get(strconv.Itoa(i))
		if err != nil {
			log.Fatal(err)
		}

		if server != newServer {
			migrate++
		}
	}

	migrateRatio := float64(migrate) / float64(keys)
	fmt.Printf("%f%%\n", migrateRatio*100)
}
