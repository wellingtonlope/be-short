package repository

import (
	"context"
	"github.com/wellingtonlope/be-short/internal/domain"
)

type Shorted interface {
	Save(ctx context.Context, shorted domain.Shorted) (domain.Shorted, error)
	GetByHash(ctx context.Context, hash string) (domain.Shorted, error)
}
