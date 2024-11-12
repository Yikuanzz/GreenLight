package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router {
	// 实例化一个 httpRouter
	router := httprouter.New()

	// 注册路由失败处理
	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	// 注册路由方法错误处理
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	// 注册路由方法
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)

	// 返回路由实例
	return router
}
