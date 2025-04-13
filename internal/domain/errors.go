package domain

import "errors"

var (
	ErrAccountNotFound    = errors.New("account not found")
	ErrDuplicateAPIKey    = errors.New("duplicate api key")
	ErrInvoiceNotFound    = errors.New("invoice not found")
	ErrUnauthorizedAccess = errors.New("unauthorized access")

	ErrInvalidAmount = errors.New("invalid amount")
	ErrInvalidStatus = errors.New("invalid status")
)
