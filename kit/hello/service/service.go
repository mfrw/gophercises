package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type serializableError struct{ error }

func (s *serializableError) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Error())
}

func newSerializeableError(text string) error {
	return &serializableError{errors.New(text)}
}

type Service struct {
}

func (s Service) Hello(ctx context.Context, firstName string, lastName string) (string, error) {
	firstName = strings.Trim(firstName, "\t\r\n")
	lastName = strings.Trim(lastName, "\t\n\n")

	if len(firstName) == 0 && len(lastName) == 0 {
		return "", newSerializeableError("missing required name information")
	}
	if len(firstName) == 0 {
		return fmt.Sprintf("Hello Mr./Ms. %s, nice to meet you. Do you have a first name?", lastName), nil
	}
	if len(lastName) == 0 {
		return fmt.Sprintf("Hello %s, nice to meet you. Do you have a last name?", firstName), nil
	}
	return fmt.Sprintf("Hello %s %s, nice to meet you.", firstName, lastName), nil
}
