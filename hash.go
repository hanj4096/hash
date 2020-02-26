package main

import (
	"flag"
	"fmt"
)

var keysPtr = flag.Int("keys", 10000000, "key number")
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

	migrate := 0
	for i := 0; i < keys; i++ {
		if hash(i, nodes) != hash(i, newNodes) {
			migrate++
		}
	}

	migrateRatio := float64(migrate) / float64(keys)
	fmt.Printf("%f%%\n", migrateRatio*100)
}
