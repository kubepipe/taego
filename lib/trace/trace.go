package trace

type Trace struct {
	Id       string `json:"id,omitempty"`
	SrcIp    string `json:"srcIp,omitempty"`
	ServerIp string `json:"serverIp,omitempty"`
}
