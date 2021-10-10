/*
Copyright 2021 Compose Generator Contributors
All rights reserved ©
*/
package terminal

type EraseLineMode int

const (
	ERASE_LINE_END EraseLineMode = iota
	ERASE_LINE_START
	ERASE_LINE_ALL
)
