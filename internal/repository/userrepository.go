package repo

import (
	"context"

	"github.com/Reshnyak/AcornsTask/internal/models"
)

//go:generate mockgen -source=userrepository.go destination=../testdata/mocks/userrepository_mock.go -package mock
type UserRepository interface {
	SaveUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, id models.UserID) (*models.User, error)
}
