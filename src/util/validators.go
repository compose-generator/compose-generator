package util

import (
	"errors"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

// PortValidator is a validator function to check if a port number is valid
func PortValidator(val interface{}) error {
	if number, err := strconv.Atoi(val.(string)); err != nil || number < 0 || number > 65535 {
		return errors.New("please provide an integer value between 0 and 65535")
	}
	return nil
}

// EnvVarNameValidator is a validator to check if a name of an environment variable is valid
func EnvVarNameValidator(val interface{}) error {
	initValidator()
	if validate.Var(strings.ReplaceAll(val.(string), "_", ""), "required,alphanum") != nil {
		return errors.New("please provide a valid name. Only alphanumeric chars and underscores are allowed")
	}
	return nil
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func initValidator() {
	if validate == nil {
		validate = validator.New()
	}
}
