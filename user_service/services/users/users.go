package users

import (
	"context"

	"github.com/cjodra14/basketball-management/user_service/api/models"
	"github.com/cjodra14/basketball-management/user_service/storage"
	_ "github.com/lib/pq"
	"go.opentelemetry.io/otel"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userStorage storage.UserStorage
}

func New(userStorage storage.UserStorage) *UserService {
	return &UserService{
		userStorage: userStorage,
	}
}

func (userService *UserService) Register(ctx context.Context, userRegister models.UserRegister) error {
	tracer := otel.Tracer("register-user-service")
	_, span := tracer.Start(ctx, "register user")
	defer span.End()

	hash, err := bcrypt.GenerateFromPassword([]byte(userRegister.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	userRegister.Password = string(hash)

	if err := userService.userStorage.Register(ctx, userRegister); err != nil {
		return err
	}

	return nil
}

func (userService *UserService) Login(ctx context.Context, userLogin models.UserLogin) error {
	tracer := otel.Tracer("login-user-service")
	_, span := tracer.Start(ctx, "login user")
	defer span.End()

	user, err := userService.userStorage.Login(ctx, userLogin)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password))
	if err != nil {
		// Password does not match
		return err
	}

	return nil
}

func (userService *UserService) Get(ctx context.Context, userID string) (models.User, error) {
	tracer := otel.Tracer("get-user-service")
	_, span := tracer.Start(ctx, "getting user")
	defer span.End()

	user, err := userService.userStorage.Get(ctx, userID)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
