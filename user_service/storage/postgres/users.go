package postgres

import (
	"context"
	"database/sql"

	"github.com/cjodra14/basketball-management/user_service/api/models"
	"github.com/cjodra14/basketball-management/user_service/configuration"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"go.opentelemetry.io/otel"
)

type PostgresUserStorage struct {
	database *sql.DB
}

func NewPostgresUserStorage(configuration configuration.PostgresConfiguration) (*PostgresUserStorage, error) {
	db, err := sql.Open("postgres", configuration.URI)
	if err != nil {
		return &PostgresUserStorage{}, err
	}

	err = db.Ping()
	if err != nil {
		return &PostgresUserStorage{}, err
	}

	return &PostgresUserStorage{
		database: db,
	}, nil
}

func (postgresUserStorage *PostgresUserStorage) Register(ctx context.Context, user models.UserRegister) error {
	tracer := otel.Tracer("register-user-postgres")
	_, span := tracer.Start(ctx, "register user")
	defer span.End()
	const query = "INSERT INTO usuarios (id, nombre, apellidos, email, password, rol) VALUES ($1, $2, $3, $4, $5, $6)"
	id := uuid.New()
	rolID := uuid.New()
	_, err := postgresUserStorage.database.Exec(query, id.String(), user.Name, user.Surname, user.Email, user.Password, rolID)
	if err != nil {
		return err
	}

	return nil
}

func (postgresUserStorage *PostgresUserStorage) Login(ctx context.Context, user models.UserLogin) (models.User, error) {
	tracer := otel.Tracer("login-user-postgres")
	_, span := tracer.Start(ctx, "login user")
	defer span.End()
	const query = "SELECT * FROM usuarios WHERE email = $1"
	foundUser := models.User{}
	err := postgresUserStorage.database.QueryRow(query, user.Email).Scan(&foundUser.ID, &foundUser.Role, &foundUser.Email, &foundUser.Password, &foundUser.FirstName, &foundUser.LastName)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, err
		}
		return models.User{}, err
	}

	return models.User{
		ID:        string(foundUser.ID),
		FirstName: foundUser.FirstName,
		LastName:  foundUser.LastName,
		Email:     foundUser.Email,
		Role:      string(foundUser.Role),
		Password:  foundUser.Password,
	}, nil
}

func (postgresUserStorage *PostgresUserStorage) Get(ctx context.Context, userID string) (models.User, error) {
	tracer := otel.Tracer("get-user-postgres")
	_, span := tracer.Start(ctx, "getting user")
	defer span.End()
	const query = "SELECT * FROM usuarios WHERE id = $1"
	user := models.User{}
	err := postgresUserStorage.database.QueryRow(query, userID).Scan(&user.ID, &user.Role, &user.Email, &user.Password, &user.FirstName, &user.LastName)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, err
		}
		return models.User{}, err
	}

	return models.User{
		ID:        string(user.ID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Role:      string(user.Role),
	}, nil
}
