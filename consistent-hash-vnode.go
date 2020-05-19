package main

import (
	"flag"
	"fmt"
	"log"
	"stathat.com/c/consistent"
	"strconv"
	"strings"
)

var keysPtr = flag.Int("keys", 10000, "key number")
var nodesPtr = flag.Int("nodes", 3, "node number of old cluster")
var vnodesPtr = flag.Int("vnodes", 100, "node number of new cluster")

func ratio(v1, v2 int) string {

	r := float64(v1) / float64(v2)
	return fmt.Sprintf("%f%%", r*100)
}

func main() {

	flag.Parse()
	var keys = *keysPtr
	var nodes = *nodesPtr
	var vnodes = *vnodesPtr
	var nodeStr = ""

	c := consistent.New()
	for i := 0; i < nodes; i++ {
		nodeStr = fmt.Sprintf("node-%d", i)
		c.Add(nodeStr)
	}

	vnodeC := consistent.New()
	for i := 0; i < nodes; i++ {
		for j := 0; j < vnodes; j++ {
			nodeStr = fmt.Sprintf("node-%d-vnode-%d", i, j)
			vnodeC.Add(nodeStr)
		}
	}

	node0, node1, node2 := 0, 0, 0
	for i := 0; i < keys; i++ {
		server, err := c.Get(strconv.Itoa(i))
		if err != nil {
			log.Fatal(err)
		}
		if strings.Compare(server, "node-0") == 0 {
			node0++
		} else if strings.Compare(server, "node-1") == 0 {
			node1++
		} else if strings.Compare(server, "node-2") == 0 {
			node2++
		} else {
			fmt.Println("unknown server:", server)
		}
	}

	fmt.Println("normal mode: node0", ratio(node0, keys/3),
		", node1", ratio(node1, keys/3),
		", node2", ratio(node2, keys/3))

	node0, node1, node2 = 0, 0, 0
	for i := 0; i < keys; i++ {
		server, err := vnodeC.Get(strconv.Itoa(i))
		if err != nil {
			log.Fatal(err)
		}
		if strings.HasPrefix(server, "node-0") {
			node0++
		} else if strings.HasPrefix(server, "node-1") {
			node1++
		} else if strings.HasPrefix(server, "node-2") {
			node2++
		} else {
			fmt.Println("unknown server:", server)
		}
	}

	fmt.Println("vnode mode: node0", ratio(node0, keys/3),
		", node1", ratio(node1, keys/3),
		", node2", ratio(node2, keys/3))

}
