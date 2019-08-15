package main

import (
	"github.com/thinkoner/thinkgo"
	"github.com/thinkoner/thinkgo/app"
	"github.com/thinkoner/thinkgo/context"
	"github.com/thinkoner/thinkgo/router"
	"thinkgo/controller"
	"thinkgo/handler"
	"thinkgo/middleware"
)

func main() {
	thinkGo := thinkgo.BootStrap()
	thinkGo.RegisterRoute(routeFunc)
	thinkGo.RegisterHandler(func(app *app.Application) app.Handler { //全局处理器
		return &handler.HandLog{}
	})

	thinkGo.Run("127.0.0.1:8081")
}

func routeFunc(route *router.Route) {
	route.Get("/", controller.Hello)
	route.Get("/ping", controller.Ping)
	// Dependency injection
	route.Get("/user/{name}", controller.UserName)
	//中间件
	route.Get("/foo", func(request *context.Request) *context.Response {
		return thinkgo.Text("Hello ThinkGo !")
	}).Middleware(middleware.CheckParam)

	//重定向
	route.Get("/redirect", func(request *context.Request) *context.Response {
		return context.Redirect("http://google.com")
	})

	route.Get("/cookie", func(request *context.Request) *context.Response {
		response := thinkgo.Text("cookies")
		response.Cookie("name", "ys")
		return response
	})

	route.Get("/tpl", func(request *context.Request) *context.Response {
		data := map[string]interface{}{"Title": "ThinkGo", "Message": "Hello ThinkGo !"}
		return thinkgo.Render("tpl.html", data)
	})

	route.Get("/getsession", controller.GetSession)
	route.Get("/setsession", controller.SetSession)
	route.Get("/info", controller.Info)
	route.Get("/setcache/{k}/{v}", controller.SetCache)
	route.Get("/getcache/{k}", controller.GetCache)
}
