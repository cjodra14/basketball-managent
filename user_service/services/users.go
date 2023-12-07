package services

import (
	"context"

	"github.com/cjodra14/basketball-management/user_service/api/models"
)

type UserService interface {
	Register(ctx context.Context, user models.UserRegister) error
	Login(ctx context.Context, user models.UserLogin) error
	Get(ctx context.Context, userID string) (models.User, error)
}
