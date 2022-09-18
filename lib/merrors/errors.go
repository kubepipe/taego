package merrors

import "taego/lib/util"

type Myerr string

func (e Myerr) Error() string {
	return string(e)
}

func (e Myerr) Code() uint32 {
	return util.HashStr(string(e))
}

func New(message string) Myerr {
	return Myerr(message)
}

func NewByError(err error) Myerr {
	return Myerr(err.Error())
}

const (
	ERROR_UNAUTHORIZED = Myerr("unauthorized")
)
