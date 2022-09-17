package mconst

import "taego/lib/trace"

type (
	Response struct {
		Message string       `json:"message"`
		Success bool         `json:"success"`
		Trace   *trace.Trace `json:"trace,omitempty"`
	}
)
