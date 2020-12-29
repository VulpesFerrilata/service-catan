package controller

import "github.com/kataras/iris/v12/websocket"

type WebsocketController interface {
}

func NewWebsocketController() WebsocketController {
	return &websocketController{}
}

type websocketController struct {
	*websocket.NSConn `stateless:"true"`
}

func (wc websocketController) Namespace() string {
	return "default"
}
