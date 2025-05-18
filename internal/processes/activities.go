package processes

import (
	"context"

	"github.com/google/uuid"

	repo "github.com/go-microfrontend/items-repository/internal/repository"
)

type Repo interface {
	GetCategories(ctx context.Context) ([]repo.Category, error)
	GetProductsByCategory(ctx context.Context, category *string) ([]repo.Product, error)
	GetProductByID(ctx context.Context, productID uuid.UUID) (repo.Product, error)
	GetProductCharacteristicByID(
		ctx context.Context,
		productID uuid.UUID,
	) (repo.ProductCharacteristic, error)
}

type Activities struct {
	repo Repo
}

func New(repo Repo) *Activities {
	return &Activities{repo: repo}
}

func (a *Activities) GetCategories(ctx context.Context) ([]repo.Category, error) {
	return a.repo.GetCategories(ctx)
}

func (a *Activities) GetProductsByCategory(
	ctx context.Context,
	category string,
) ([]repo.Product, error) {
	return a.repo.GetProductsByCategory(ctx, &category)
}

func (a *Activities) GetProductByID(
	ctx context.Context,
	productID string,
) (repo.Product, error) {
	uuid, _ := uuid.Parse(productID)
	return a.repo.GetProductByID(ctx, uuid)
}

func (a *Activities) GetProductCharacteristicByID(
	ctx context.Context,
	productID string,
) (repo.ProductCharacteristic, error) {
	uuid, _ := uuid.Parse(productID)
	return a.repo.GetProductCharacteristicByID(ctx, uuid)
}
