package manager

import (
	"context"
	socketio "github.com/googollee/go-socket.io"
	"log"
	"t-board/internal/controller/ws/events"
	"t-board/internal/controller/ws/types"
	"t-board/internal/entity"
	"t-board/internal/usecase"
)

type EventManager interface {
	PopulateEvents()
}

type eventManager struct {
	so           *socketio.Server
	userUseCase  usecase.UserUseCase
	boardUseCase usecase.BoardUseCase
}

func CreateEventManager(so *socketio.Server, u usecase.UserUseCase, b usecase.BoardUseCase) EventManager {
	return &eventManager{so, u, b}
}

func (m *eventManager) PopulateEvents() {
	m.appendErrorHandler()

	m.appendLoginEvent()
	m.appendRegisterEvent()
}

func (m *eventManager) appendErrorHandler() {
	m.so.OnError("/", func(s socketio.Conn, e error) {
		log.Println("Error:", e)
	})
}

func (m *eventManager) appendLoginEvent() {
	m.so.OnEvent("/", events.LoginRequest, func(s socketio.Conn, r types.LoginUserRequest) {
		token, err := m.userUseCase.Login(context.Background(), r.Email, r.Password)
		if err != nil {
			s.Emit(events.Error, types.ErrorResponse{Error: err.Error()})
			return
		}

		s.Emit(events.LoginResponse, types.LoginUserResponse{Token: token})
	})
}

func (m *eventManager) appendRegisterEvent() {
	m.so.OnEvent("/", events.RegisterRequest, func(s socketio.Conn, r types.CreateUser) {
		registerUser := &entity.User{
			Name:     r.Name,
			Email:    r.Email,
			Password: r.Password,
		}
		if err := m.userUseCase.Register(context.Background(), registerUser); err != nil {
			s.Emit(events.Error, types.ErrorResponse{Error: err.Error()})
		}
	})
}
