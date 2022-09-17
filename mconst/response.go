package mconst

type (
	Response struct {
		Message string `json:"message"`
		Success bool   `json:"success"`
		Trace   *Trace `json:"trace,omitempty"`
	}

	Trace struct {
		Id       string `json:"id,omitempty"`
		SrcIp    string `json:"srcIp,omitempty"`
		ServerIp string `json:"serverIp,omitempty"`
	}
)
