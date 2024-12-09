package storage

import (
	"context"

	"github.com/Reshnyak/AcornsTask/internal/configs"
	"github.com/Reshnyak/AcornsTask/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStore struct {
	config *configs.TaskStore
	client *mongo.Client
}

func NewTaskStore(ctx context.Context, cfgStore *configs.TaskStore) (*MongoStore, error) {

	msClient, err := mongo.Connect(ctx, options.Client().ApplyURI(cfgStore.Url))
	if err != nil {
		return nil, err
	}
	mongoStore := &MongoStore{
		config: cfgStore,
		client: msClient,
	}
	return mongoStore, nil
}

func (ms *MongoStore) AddTask(ctx context.Context, task *models.Task) error {
	return nil
}
func (ms *MongoStore) GetTask(ctx context.Context, id models.Task) (*models.Task, error) {
	return nil, nil
}
func (ms *MongoStore) UpdateTask(ctx context.Context, task *models.Task) error {
	return nil
}
