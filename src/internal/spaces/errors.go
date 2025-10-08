package spaces

import "errors"

var (
	ErrUserNotActive   = errors.New("user is not active")
	ErrUserNotVerified = errors.New("user email is not verified")

	ErrSlugTooLong        = errors.New("slug exceeds maximum length of 20 characters")
	ErrSlugTooShort       = errors.New("slug must be at least 3 characters long")
	ErrTitleTooLong       = errors.New("title exceeds maximum length of 100 characters")
	ErrTitleTooShort      = errors.New("title must be at least 3 characters long")
	ErrDescriptionTooLong = errors.New("description exceeds maximum length of 500 characters")
)
