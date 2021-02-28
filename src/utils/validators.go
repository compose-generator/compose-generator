package utils

import (
	"errors"
	"strconv"
)

// PortValidator is a validator function to check if a port number is valid
func PortValidator(val interface{}) error {
	if number, err := strconv.Atoi(val.(string)); err != nil || number < 0 || number > 65535 {
		return errors.New("Please provide an integer value between 0 and 65535")
	}
	return nil
}
