package network

import (
	"gochat/types"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"golang.org/x/text/cases"
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

	Join  chan *client // Socket이 연결되었을 때 작동
	Leave chan *client // Socket이 끊어졌을 때 작동

	Clients map[*client]bool // 방에 있는 클라이언트 정보를 저장
}

type client struct {
	Send   chan *message
	Room   *Room
	Name   string
	Socket *websocket.Conn
}

func NewRoom() *Room {
	return &Room{
		Forward: make(chan *message),
		Join:    make(chan *client),
		Leave:   make(chan *client),
		Clients: make(map[*client]bool),
	}
}

func (r *Room) RunInit() {
	// Room에 있는 모든 채널 값을 받는 역할
	for {
		select {
		case client := <-r.Join:
			r.Clients[client] = true
		case client := <-r.Leave:
			r.Clients[client] = false
			close(client.Send)
			delete(r.Clients, client)
		case msg := <-r.Forward:
			for client, v := r.Clients {

			}
		}
	}
}

func (r *Room) SocketServe(c *gin.Context) {
	socket, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		panic(err)
	}

	userCookie, err := c.Request.Cookie("auth")
	if err != nil {
		panic(err)
	}

	client := &client{
		Socket: socket,
		Send: make(chan *message, types.MessageBufferSize),
		Room: r,
		Name: userCookie.Value
	}

	// 진입
	r.Join <- client

	// 퇴장
	defer func() { r.Leave <- client }
}
