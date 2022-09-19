package merrors

type Merr interface {
	Error() string
	Code() int
}
