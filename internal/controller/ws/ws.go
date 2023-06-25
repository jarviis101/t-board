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
)

type ws struct {
}

func CreateWSServer() controller.Server {
	return &ws{}
}

func (s *ws) RunServer() error {
	so := socketio.NewServer(&engineio.Options{
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

	so.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Println("connected:", s.ID())

		return nil
	})
	type Message struct {
		Msg string `json:"msg"`
	}
	so.OnEvent("/", "notice", func(s socketio.Conn, msg Message) {
		log.Println("notice:", msg)
		s.Emit("reply", "have "+msg.Msg)
	})

	so.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		if err := s.Close(); err != nil {
			return ""
		}
		return last
	})

	so.OnError("/", func(s socketio.Conn, e error) {
		log.Println("meet error:", e)
	})

	so.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("closed", reason)
	})

	go func() {
		if err := so.Serve(); err != nil {
			log.Fatalf("Error: %s", err.Error())
		}
	}()

	http.Handle("/socket.io/", so)
	log.Fatal(http.ListenAndServe(":5000", nil))

	return nil
}
