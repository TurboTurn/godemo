package handler

import (
	"github.com/thinkoner/thinkgo/app"
	"github.com/thinkoner/thinkgo/context"
	"github.com/thinkoner/thinkgo/log"
	"time"
)

type HandLog struct {
}

func (h *HandLog) Process(r *context.Request, next app.Closure) interface{} {
	log.Info("[%s] %q %v", r.Method(), r.Path(), time.Now())
	return next(r)
}
