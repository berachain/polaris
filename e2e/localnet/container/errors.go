package container

import "errors"

// Errors returned by the localnet package.
var (
	ErrEmptyName        = errors.New("name cannot be empty")
	ErrEmptyImageName   = errors.New("image name cannot be empty")
	ErrEmptyContext     = errors.New("context cannot be empty")
	ErrEmptyDockerfile  = errors.New("dockerfile cannot be empty")
	ErrEmptyHTTPAddress = errors.New("http address cannot be empty")
	ErrEmptyWSAddress   = errors.New("ws address cannot be empty")
)
