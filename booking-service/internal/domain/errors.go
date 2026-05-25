package domain

import "errors"

var (
	ErrUserAlreadyExists   = errors.New("user already exists")
	ErrInvalidCredentials  = errors.New("invalid credentials")
	ErrHouseNotFound       = errors.New("house not found")
	ErrInvalidDateFormat   = errors.New("invalid date format")
	ErrInvalidBookingDates = errors.New("date_end must be after date_start")
)
