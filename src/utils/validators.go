package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"unicode/utf8"
)

// NoSpacesValidator is a validator to check if a string contains space chars
func NoSpacesValidator(val interface{}) error {
	// TODO: Implement validator
	return nil
}

// EmailValidator is a validator to check if an email is valid
func EmailValidator(val interface{}) error {
	var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	email := val.(string)
	if len(email) < 3 || len(email) > 254 || !emailRegex.MatchString(email) {
		return errors.New("please provide a valid email address")
	}
	return nil
}

// PortValidator is a validator function to check if a port number is valid
func PortValidator(val interface{}) error {
	if number, err := strconv.Atoi(val.(string)); err != nil || number < 0 || number > 65535 {
		return errors.New("please provide an integer value between 0 and 65535")
	}
	return nil
}

// HostNameValidator is a validator to check if a hostname is valid.
func HostNameValidator(val interface{}) error { // Thanks to @chmike for this code snippet
	name := val.(string)
	switch {
	case len(name) == 0:
		return nil
	case len(name) > 255:
		return fmt.Errorf("name length is %d, can't exceed 255", len(name))
	}
	var l int
	for i := 0; i < len(name); i++ {
		b := name[i]
		if b == '.' {
			switch {
			case i == l:
				return fmt.Errorf("invalid character '%c' at offset %d: label can't begin with a period", b, i)
			case i-l > 63:
				return fmt.Errorf("byte length of label '%s' is %d, can't exceed 63", name[l:i], i-l)
			case name[l] == '-':
				return fmt.Errorf("label '%s' at offset %d begins with a hyphen", name[l:i], l)
			case name[i-1] == '-':
				return fmt.Errorf("label '%s' at offset %d ends with a hyphen", name[l:i], l)
			}
			l = i + 1
			continue
		}
		if !(b >= 'a' && b <= 'z' || b >= '0' && b <= '9' || b == '-' || b >= 'A' && b <= 'Z') {
			c, _ := utf8.DecodeRuneInString(name[i:])
			if c == utf8.RuneError {
				return fmt.Errorf("invalid rune at offset %d", i)
			}
			return fmt.Errorf("invalid character '%c' at offset %d", c, i)
		}
	}

	switch {
	case l == len(name):
		return fmt.Errorf("missing top level domain, domain can't end with a period")
	case len(name)-l > 63:
		return fmt.Errorf("byte length of top level domain '%s' is %d, can't exceed 63", name[l:], len(name)-l)
	case name[l] == '-':
		return fmt.Errorf("top level domain '%s' at offset %d begins with a hyphen", name[l:], l)
	case name[len(name)-1] == '-':
		return fmt.Errorf("top level domain '%s' at offset %d ends with a hyphen", name[l:], l)
	case name[l] >= '0' && name[l] <= '9':
		return fmt.Errorf("top level domain '%s' at offset %d begins with a digit", name[l:], l)
	}
	return nil
}

// EnvVarNameValidator is a validator to check if a name of an environment variable is valid
func EnvVarNameValidator(val interface{}) error {
	for _, char := range val.(string) {
		if (char < 'a' || char > 'z') && (char < 'A' || char > 'Z') && (char < '0' || char > '9') && char != '_' {
			return errors.New("please provide a valid name. Only alphanumeric chars and underscores are allowed")
		}
	}
	return nil
}
