package repo

import (
	"context"

	"github.com/Reshnyak/AcornsTask/internal/models"
)

type TaskRepository interface {
	AddTask(ctx context.Context, task *models.Task) error
	GetTask(ctx context.Context, id models.Task) (*models.Task, error)
	UpdateTask(ctx context.Context, task *models.Task) error
}
