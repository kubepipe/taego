package mconst

type myerr string

func (e myerr) Error() string {
	return string(e)
}

const (
	ERROR_UNAUTHORIZED = myerr("unauthorized")
)
