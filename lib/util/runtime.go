package util

import "runtime"

func GetFuncName() string {
	if pc, _, _, ok := runtime.Caller(1); ok {
		return runtime.FuncForPC(pc).Name()
	}
	return "unknowfunc"
}
