package services

import "errors"

var (
  ErrUserNotFound = errors.New("User not found")
  ErrAlreadyVerified = errors.New("Email already verified")
  ErrTooManyRequest = errors.New("Too many requests, try again later")
  ErrInternalServer = errors.New("Internal server error")
  ErrEmailFailed = errors.New("Failed to send email")
  ErrNotVerified = errors.New("Email not verified")
  ErrOauthUser = errors.New("Cannot reset password for oauth user")
)
