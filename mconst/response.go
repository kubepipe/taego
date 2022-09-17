package mconst

type (
	Response struct {
		Message string     `json:"message,omitempty"`
		Success bool       `json:"success,omitempty"`
		Trace   *TraceInfo `json:"trace,omitempty"`
	}

	TraceInfo struct {
		Id       int32  `json:"id,omitempty"`
		SourceIp string `json:"sourceIp,omitempty"`
		ServerIp string `json:"serverIp,omitempty"`
	}
)
