/*
Copyright 2021 Compose Generator Contributors
All rights reserved ©
*/
package model

// CComDataInput represents the structure, in which data gets passed to CCom
type CComDataInput struct {
	Services SelectedTemplates `json:"services,omitempty"`
	Var      map[string]string `json:"var,omitempty"`
}
