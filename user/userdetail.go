// Copyright (C) 2019-2021, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package user

type Details interface {
	GetLimit(url string) int

	GetPassword() string
	GetUsername() string

	IsAccountNonExpired() bool
	IsAccountNonLocked() bool
	IsCredentialsNonExpired() bool

	IsEnabled() bool
}
