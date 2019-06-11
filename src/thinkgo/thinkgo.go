package main

import (
	"fmt"
	"github.com/thinkoner/thinkgo"
	"github.com/thinkoner/thinkgo/context"
	"github.com/thinkoner/thinkgo/router"
)

func main() {
	app := thinkgo.BootStrap()
	/*app.RegisterRoute(func(route *router.Route) {

		route.Get("/", func(req *context.Request) *context.Response {
			return thinkgo.Text("Hello ThinkGo !")
		})

		route.Get("/ping", func(req *context.Request) *context.Response {
			return thinkgo.Json(map[string]string{
				"message": "pong",
			})
		})

		// Dependency injection
		route.Get("/user/{name}", func(req *context.Request, name string) *context.Response {
			return thinkgo.Text(fmt.Sprintf("Hello %s !", name))
		})
	})*/
	app.RegisterRoute(routeFun)
	// listen and serve on 0.0.0.0:9011
	app.Run("127.0.0.1:8081")
}

/* 将写在参数的函数独立出来*/
func routeFun(route *router.Route) {
	route.Get("/", func(req *context.Request) *context.Response {
		return thinkgo.Text("Hello ThinkGo !")
	})

	route.Get("/ping", func(req *context.Request) *context.Response {
		return thinkgo.Json(map[string]string{
			"message": "pong",
		})
	})

	// Dependency injection
	route.Get("/user/{name}", func(req *context.Request, name string) *context.Response {
		return thinkgo.Text(fmt.Sprintf("Hello %s !", name))
	})
}
