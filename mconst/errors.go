package mconst

type myerr string

func (e myerr) Error() string {
	return string(e)
}

const (
	EXAMPLE_ERROR = myerr("some error")
)
