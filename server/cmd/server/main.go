package main

import (
	"log"
	"net/http"

	"github.com/TarekkMA/GeoChat/app/wsapp"
	"github.com/TarekkMA/GeoChat/pkg/ws"
	"github.com/gorilla/mux"
)

func main() {
	log.Println("WebServer is starting")

	wsServer := getWebSocketServer()

	r := mux.NewRouter()

	r.Handle("/ws", wsServer.Run())

	log.Fatalln(http.ListenAndServe(":8080", r))
}

func getWebSocketServer() *ws.Server {
	pool := ws.NewPool()
	handler := wsapp.NewDelegate(pool)
	pool.SetDelegate(handler)
	return ws.NewServer(pool)
}
