package utils

import (
	"bytes"
	"errors"
)

func RaiseError(message string, err error) error {
	var buffer bytes.Buffer

	buffer.WriteString(message)
	buffer.WriteString(err.Error())
	return errors.New(buffer.String())
}