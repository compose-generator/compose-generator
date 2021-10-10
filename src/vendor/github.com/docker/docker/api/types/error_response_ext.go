/*
Copyright 2021 Compose Generator Contributors
All rights reserved ©
*/
package types

// Error returns the error message
func (e ErrorResponse) Error() string {
	return e.Message
}
