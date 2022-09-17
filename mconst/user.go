package mconst

type Userinfo struct {
	Name   string `json:"name,omitempty"`
	Erp    string `json:"erp,omitempty"`
	Org    string `json:"org,omitempty"`
	Tenant string `json:"tenant,omitempty"`
}
