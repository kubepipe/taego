package mconst

type (
	Response struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
		Success bool   `json:"success"`
		Trace   *Trace `json:"trace,omitempty"`
	}

	Trace struct {
		Id    string `json:"id"`
		SrcIp string `json:"srcIp"`
	}
)

type (
	ResOk struct {
		Data string `json:"data,omitempty"`
	}
)

type (
	ResCreateTask struct {
		VersionId int `json:"resourceVersionId"`
	}
)
