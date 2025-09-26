package service

import "fmt"

// UserService handles user-related operations.
type UserService struct{}

// GetUser returns a dummy user string.
func (s UserService) GetUser(id string) string {
	return fmt.Sprintf("User ID: %s", id)
}
