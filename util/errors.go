package util

import (
	"fmt"
)

// Error represents a structured error with a code and a message
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Error implements the error interface for CustomError
func (e *Error) Error() string {
	return fmt.Sprintf("Error %s: %s", e.Code, e.Message)
}

// NewError creates a new instance of CustomError
func NewError(code, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

// ValidationError used in general, when appeared error doesn't really fit into any other defined errors
func ValidationError(message string) *Error {
	return NewError("001", message)
}

// ExpiredCardError used when card date is expired
func ExpiredCardError() *Error {
	return NewError("002", "Card is expired")
}

// InvalidCardNumberError used when card number doesn't pass luhn algorithm, which means it incorrect
func InvalidCardNumberError() *Error {
	return NewError("003", "Card number is invalid")
}

// InvalidExpirationDateError used when invalid data passed about the card expiration date
func InvalidExpirationDateError() *Error {
	return NewError("004", "Expiration date is invalid")
}

func InvalidCardNumberAndExpirationDateError() *Error {
	return NewError("005", "Card Number and Expiration date is invalid")
}
