package entities

import "github.com/google/uuid"

type User struct {
	UUID     uuid.UUID
	Username string
	Password string
	Email    string
	Role     string
}
