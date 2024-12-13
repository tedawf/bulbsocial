package auth

import "time"

// TokenMaker is an interface for managing tokens
type TokenMaker interface {
	// CreateToken creates a new token for a specific username and duration
	CreateToken(userID int64, duration time.Duration) (string, error)

	// VerifyToken checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
