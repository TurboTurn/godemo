package middleware

import (
	"github.com/thinkoner/thinkgo"
	"github.com/thinkoner/thinkgo/context"
	"github.com/thinkoner/thinkgo/router"
)

func CheckParam(request *context.Request, next router.Closure) interface{} {
	if _, err := request.Input("name"); err != nil {
		return thinkgo.Text("Invalid parameters")
	}
	return next(request)
}
