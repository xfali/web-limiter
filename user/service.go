// Copyright (C) 2019-2021, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package user

type Service interface {
	LoadUserByUsername(username string) (Details, error)
}
