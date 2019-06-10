package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
)

type controllerInfo struct {
	url            string
	controllerType reflect.Type
}

type ControllerRegistor struct {
	routers []*controllerInfo
}

type ControllerInterface interface {
	Do()
}

type UserController struct {
}

type DefaultController struct {
}

func (u *UserController) Do() {
	fmt.Println("I`m UserController")
}

func (d *DefaultController) Do() {
	fmt.Println("I`m DefaultController")
}

func (p *ControllerRegistor) Add(pattern string, c ControllerInterface) {

	//now create the Route
	t := reflect.TypeOf(c).Elem()
	route := &controllerInfo{}
	route.url = pattern
	route.controllerType = t
	p.routers = append(p.routers, route)

}

// AutoRoute
func (p *ControllerRegistor) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var started bool
	requestPath := r.URL.Path

	fmt.Println(requestPath)

	//find a matching Route
	for _, route := range p.routers {

		if requestPath == route.url {
			vc := reflect.New(route.controllerType)
			method := vc.MethodByName("Do")
			method.Call(nil)
			started = true
			fmt.Fprintf(w, "Hello "+route.controllerType.Name())
			break
		}
	}

	//if no matches to url, throw a not found exception
	if started == false {
		http.NotFound(w, r)
	}
}

//实现简单的路由功能
func main() {
	mux := &ControllerRegistor{}

	mux.Add("/", &DefaultController{})
	mux.Add("/user", &UserController{})

	err := http.ListenAndServe(":9527", mux)
	if err != nil {
		_, err := fmt.Fprintln(os.Stderr, "发生错误！")
		fmt.Println(err)
	}
}
