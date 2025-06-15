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

type Cache interface {
	GetCategories(ctx context.Context) ([]repo.Category, error)
	GetProductsByCategory(ctx context.Context, category *string) ([]repo.Product, error)
	GetProductByID(ctx context.Context, productID uuid.UUID) (repo.Product, error)
	GetProductCharacteristicByID(
		ctx context.Context,
		productID uuid.UUID,
	) (repo.ProductCharacteristic, error)
}

type Activities struct {
	repo  Repo
	cache Cache
}

func New(repo Repo, cache Cache) *Activities {
	return &Activities{repo: repo, cache: cache}
}

func (a *Activities) GetCategories(ctx context.Context) ([]repo.Category, error) {
	categories, err := a.cache.GetCategories(ctx)
	if err == nil {
		return categories, nil
	}

	categories, err = a.repo.GetCategories(ctx)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (a *Activities) GetProductsByCategory(
	ctx context.Context,
	category string,
) ([]repo.Product, error) {
	products, err := a.cache.GetProductsByCategory(ctx, &category)
	if err == nil {
		return products, nil
	}

	products, err = a.repo.GetProductsByCategory(ctx, &category)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (a *Activities) GetProductByID(
	ctx context.Context,
	productID string,
) (repo.Product, error) {
	uuid, _ := uuid.Parse(productID)

	product, err := a.cache.GetProductByID(ctx, uuid)
	if err == nil {
		return product, nil
	}

	return a.repo.GetProductByID(ctx, uuid)
}

func (a *Activities) GetProductCharacteristicByID(
	ctx context.Context,
	productID string,
) (repo.ProductCharacteristic, error) {
	uuid, _ := uuid.Parse(productID)

	char, err := a.cache.GetProductCharacteristicByID(ctx, uuid)
	if err == nil {
		return char, nil
	}
	return a.repo.GetProductCharacteristicByID(ctx, uuid)
}
