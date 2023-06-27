package ws

import (
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	"log"
	"net/http"
	"t-board/internal/controller"
	"t-board/internal/controller/ws/manager"
	"t-board/internal/usecase"
)

type ws struct {
	userUseCase  usecase.UserUseCase
	boardUseCase usecase.BoardUseCase
}

func CreateWSServer(u usecase.UserUseCase, b usecase.BoardUseCase) controller.Server {
	return &ws{u, b}
}

func (s *ws) RunServer() error {
	so := s.createServer()

	eventManager := manager.CreateEventManager(so, s.userUseCase, s.boardUseCase)
	eventManager.PopulateEvents()

	go func() {
		if err := so.Serve(); err != nil {
			log.Printf("Error: %s\n", err.Error())
		}
	}()

	http.Handle("/socket.io/", so)

	return http.ListenAndServe(":5000", nil)
}

func (s *ws) createServer() *socketio.Server {
	return socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			&polling.Transport{
				CheckOrigin: func(r *http.Request) bool {
					return true
				},
			},
			&websocket.Transport{
				CheckOrigin: func(r *http.Request) bool {
					return true
				},
			},
		},
	})
}
