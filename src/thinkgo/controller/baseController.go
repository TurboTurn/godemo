package controller

import (
	"fmt"
	"github.com/thinkoner/thinkgo"
	"github.com/thinkoner/thinkgo/context"
	"github.com/thinkoner/thinkgo/log"
	"thinkgo/store"
	"time"
)

func Hello(request *context.Request) *context.Response {
	data := map[string]interface{}{
		"Title":   "ThinkGo",
		"Message": "Hello ThinkGo !",
		"name":    "yishuang",
	}
	return thinkgo.Render("index.html", data)
}
func Ping(request *context.Request) *context.Response {
	return thinkgo.Json(map[string]string{
		"message": "pong",
	})
}
func UserName(request *context.Request, name string) *context.Response {
	return thinkgo.Text(fmt.Sprintf("Hello %s !", name))
}

//todo session的存取有错误
func SetSession(request *context.Request) *context.Response {
	request.Session().Set("user", "alice")
	fmt.Println("setsession")
	return thinkgo.Text("session set")
}
func GetSession(request *context.Request) *context.Response {
	user := request.Session().Get("user")
	fmt.Println(user)
	return thinkgo.Text("session get")
}

func Info(request *context.Request) *context.Response {
	log.Debug("log")
	log.Info("info")
	log.Error("error")
	log.Alert("alert")
	return thinkgo.Text("info")
}

//var c = store.MemRepository()
var c = store.RedisReposotory()

func SetCache(request *context.Request, k string, v string) *context.Response {
	c.Put(k, v, time.Hour*24)
	return thinkgo.Text("setCache " + k + " " + v)
}
func GetCache(request *context.Request, k string) *context.Response {
	var v string
	c.Get(k, &v)
	return thinkgo.Text("getCache " + k + " " + v)
}
