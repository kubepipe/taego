package merrors

type Code int

const (
	ERROR_UNAUTHORIZED Code = iota + 10000
	ERROR_UNHEALTHY

	// TODO add new error code here

)

var errmap = map[string]Code{
	// TODO add new error descriptions here

	"unauthorized": ERROR_UNAUTHORIZED,
	"unhealthy":    ERROR_UNHEALTHY,
}

var codemap map[Code]string

func init() {
	codemap = make(map[Code]string, len(errmap))
	for k, v := range errmap {
		codemap[v] = k
	}
}

type merr string

func (e merr) Error() string {
	return string(e)
}

func (e merr) Code() int {
	if code, ok := errmap[string(e)]; ok {
		return int(code)
	}
	// normal error
	return -1
}

func New(err error) Merr {
	return merr(err.Error())
}

func Get(errcode Code) Merr {
	if errmsg, ok := codemap[errcode]; ok {
		return merr(errmsg)
	}
	return merr("unknow error")
}
