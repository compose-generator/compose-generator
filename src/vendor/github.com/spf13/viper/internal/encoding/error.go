/*
Copyright 2021 Compose Generator Contributors
All rights reserved ©
*/
package encoding

type encodingError string

func (e encodingError) Error() string {
	return string(e)
}
