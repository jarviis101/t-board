package ws

import (
	socketio "github.com/googollee/go-socket.io"
	"log"
	"net/http"
	"t-board/internal/controller"
)

type ws struct {
}

func CreateWSServer() controller.Server {
	return &ws{}
}

func (s *ws) RunServer() error {
	server := socketio.NewServer(nil)

	http.Handle("/socket.io/", server)

	log.Fatal(http.ListenAndServe(":5000", nil))

	return nil
}
