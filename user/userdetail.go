// Copyright (C) 2019-2021, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package user

import "golang.org/x/time/rate"

type Details interface {
	// 获得用户限制
	// limit：每秒访问次数。burst：最大预留访问次数。
	GetLimit(url string) *rate.Limiter

	GetUsername() string
}
