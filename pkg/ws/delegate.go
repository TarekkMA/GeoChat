package ws

import (
	"net/http"
)

type WebSocketDelegate interface {
	HandleMessage(*Client, []byte)
	HandleRegister(*Client, *http.Request) bool
}
