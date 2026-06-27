package repository

import (
	"context"

	"github.com/hcivekhsim/links-hub/models"
)

type LinkRepository interface {
	GetAll(ctx context.Context) ([]models.Link, error)
	AddLink(ctx context.Context, link models.Link) (int, error)
	ChangeLink(ctx context.Context, id int, link models.LinkUpdate) error
	RemoveLink(ctx context.Context, id int) error
}
