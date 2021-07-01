// Copyright (C) 2019-2021, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package ginlimiter

import (
	"github.com/gin-gonic/gin"
	"github.com/xfali/web-limiter/user"
	"golang.org/x/time/rate"
	"math"
	"sync"
)

const (
	Unlimited = -1
	MIN       = 0.000001
)

type UserService interface {
	LoadUserByUsername(ctx *gin.Context) (user.Details, error)
}

type LimitHandler interface {
	OnLimited(ctx *gin.Context)
}

type Filter struct {
	service  UserService
	handler  LimitHandler
	limiters sync.Map
}

func NewFilter(service UserService, handler LimitHandler) *Filter {
	ret := &Filter{
		service: service,
		handler: handler,
	}
	return ret
}

func (f *Filter) FilterHandler(ctx *gin.Context) {
	d, err := f.service.LoadUserByUsername(ctx)
	if err != nil || !f.checkLimit(ctx.Request.RequestURI, d) {
		f.handler.OnLimited(ctx)
	} else {
		ctx.Next()
	}
}

func (f *Filter) checkLimit(url string, user user.Details) bool {
	limiter := user.GetLimit(url)
	if limiter == nil {
		return true
	} else {
		return limiter.Allow()
	}
}

func compareLimiter(l *rate.Limiter, limit int) bool {
	return math.Dim(float64(l.Limit()), float64(limit)) < MIN
}

func (f *Filter) getLimiter(user user.Details) *rate.Limiter {
	if v, ok := f.limiters.Load(user.GetUsername()); ok {
		return v.(*rate.Limiter)
	} else {
		return nil
	}
}

func (f *Filter) resetLimiter(user user.Details, limit int) *rate.Limiter {
	limiter := rate.NewLimiter(rate.Limit(limit), limit)
	f.limiters.Store(user.GetUsername(), limiter)
	return limiter
}

//type AbstractUserService struct {
//}
//
//func (s *AbstractUserService) LoadUserByUsername(ctx *gin.Context) (user.Details, error) {
//
//}
//
//func (s *AbstractUserService)() {
//	limiter := f.getLimiter(user)
//	if limiter == nil || compareLimiter(limiter, limit) {
//		limiter = f.resetLimiter(user, limit)
//	}
//}
