package network

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Network struct {
	engine *gin.Engine
}

func NewServer() *Network {
	n := &Network{engine: gin.New()}

	return n
}

func (n *Network) StartServer() error {
	log.Println("Starting Server")

	return n.engine.Run(":8080")
}
