package mconst

type (
	Response struct {
		Code    uint32     `json:"code"`
		Message string     `json:"message,omitempty"`
		Success bool       `json:"success"`
		Trace   *TraceInfo `json:"trace,omitempty"`
	}

	TraceInfo struct {
		Id       int32  `json:"id,omitempty"`
		SourceIp string `json:"sourceIp,omitempty"`
		ServerIp string `json:"serverIp,omitempty"`
	}
)
