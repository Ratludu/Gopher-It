package main

import (
	"time"

	"github.com/Ratludu/Gopher-It/internal/database"
)

type User struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"apikey"`
}

func databaseUserToUser(dbUser database.User) (User, error) {
	createdAt, err := time.Parse(time.RFC3339, dbUser.CreatedAt)
	if err != nil {
		return User{}, err
	}

	updatedAt, err := time.Parse(time.RFC3339, dbUser.UpdatedAt)
	if err != nil {
		return User{}, err
	}

	return User{
		ID:        dbUser.ID,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Name:      dbUser.Name,
		ApiKey:    dbUser.ApiKey,
	}, nil
}
