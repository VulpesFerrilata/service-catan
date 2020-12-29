package router

import (
	"database/sql"

	"github.com/VulpesFerrilata/catan/infrastructure/iris/controller"

	"github.com/VulpesFerrilata/library/pkg/middleware"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/websocket"
)

type Router interface {
	InitRoutes(app *iris.Application)
}

func NewRouter(websocketController controller.WebsocketController,
	transactionMiddleware *middleware.TransactionMiddleware,
	translatorMiddleware *middleware.TranslatorMiddleware,
	errorMiddleware *middleware.ErrorMiddleware) Router {
	return &router{
		websocketController:   websocketController,
		transactionMiddleware: transactionMiddleware,
		translatorMiddleware:  translatorMiddleware,
		errorMiddleware:       errorMiddleware,
	}
}

type router struct {
	websocketController   controller.WebsocketController
	transactionMiddleware *middleware.TransactionMiddleware
	translatorMiddleware  *middleware.TranslatorMiddleware
	errorMiddleware       *middleware.ErrorMiddleware
}

func (r router) InitRoutes(app *iris.Application) {
	apiRoot := app.Party("/api")
	apiRoot.Use(
		r.transactionMiddleware.ServeWithTxOptions(&sql.TxOptions{}),
		r.translatorMiddleware.Serve,
	)

	catanRoot := apiRoot.Party("/catan")

	mvcApp := mvc.New(catanRoot)
	mvcApp.HandleWebsocket(r.websocketController)
	wsServer := websocket.New(websocket.DefaultGorillaUpgrader, mvcApp)
	catanRoot.Get("/ws", websocket.Handler(wsServer))
}
