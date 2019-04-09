package ws

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Server struct {
	pool *Pool
}

func NewServer(Pool *Pool) *Server {
	return &Server{
		pool: Pool,
	}
}

func (s *Server) Run() http.HandlerFunc {
	go s.pool.listen()
	return func(w http.ResponseWriter, r *http.Request) {
		upgrader.CheckOrigin = func(r *http.Request) bool {
			return true
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}

		client := newClient(uuid.New().String(), s.pool, conn)
		if s.pool.delegate.HandleRegister(client, r) {
			client.Pool.register <- client
			client.StartListining()
		} else {
			conn.Close()
		}
	}
}
