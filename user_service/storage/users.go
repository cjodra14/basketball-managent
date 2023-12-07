package storage

import (
	"context"

	"github.com/cjodra14/basketball-management/user_service/api/models"
)

type UserStorage interface {
	Register(ctx context.Context, user models.UserRegister) error
	Login(ctx context.Context, user models.UserLogin) (models.User, error)
	Get(ctx context.Context, userID string) (models.User, error)
}

type User struct {
	ID        []byte `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	Picture   string `json:"picture,omitempty"`
	Role      []byte `json:"role,omitempty"`
}
