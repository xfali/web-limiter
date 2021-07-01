// Copyright (C) 2019-2021, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package user

type Observer interface {
	OnUserUpdate(d Details)
	OnUserDisable(d Details)
}

type Service interface {
	LoadUserByUsername(username string) (Details, error)

	AddUserObserver(l Observer)
}
