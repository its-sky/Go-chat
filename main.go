package main

import "gochat/network"

func main() {
	n := network.NewServer()
	n.StartServer()
}
