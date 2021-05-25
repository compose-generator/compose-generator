package model

// CComDataInput represents the structure, in which data gets passed to CCom
type CComDataInput struct {
	Services map[string][]ServiceTemplateConfig `json:"services"`
	Var      map[string]string                  `json:"var"`
}
