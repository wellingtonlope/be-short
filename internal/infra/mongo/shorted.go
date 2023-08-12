package mongo

import (
	"context"
	"github.com/wellingtonlope/be-short/internal/app/repository"
	"github.com/wellingtonlope/be-short/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type shorted struct {
	collection *mongo.Collection
}

func NewShorted(database *mongo.Database) (repository.Shorted, error) {
	return &shorted{collection: database.Collection("shorted")}, nil
}

func (r shorted) Save(ctx context.Context, shorted domain.Shorted) (domain.Shorted, error) {
	//TODO implement me
	panic("implement me")
}

func (r shorted) GetByHash(ctx context.Context, hash string) (domain.Shorted, error) {
	//TODO implement me
	panic("implement me")
}
