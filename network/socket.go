package network

import (
	"gochat/types"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// HTTP 요청을 웹소켓으로 업그레이드
var upgrader = &websocket.Upgrader{
	ReadBufferSize:  types.SocketBufferSize,
	WriteBufferSize: types.MessageBufferSize,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type message struct {
	Name    string
	Message string
	Time    int64
}

type Room struct {
	Forward chan *message // 수신되는 메시지를 보관
	// 들어오는 메시지를 다른 클라이언트들에게 전달

	Join  chan *Client // Socket이 연결되었을 때 작동
	Leave chan *Client // Socket이 끊어졌을 때 작동

	Clients map[*Client]bool // 방에 있는 클라이언트 정보를 저장
}

type Client struct {
	Send   chan *message
	Room   *Room
	Name   string
	Socket *websocket.Conn
}

func NewRoom() *Room {
	return &Room{
		Forward: make(chan *message),
		Join:    make(chan *Client),
		Leave:   make(chan *Client),
		Clients: make(map[*Client]bool),
	}
}

func (r *Room) SocketServe(c *gin.Context) {
	socket, err = upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		panic(err)
	}
}
