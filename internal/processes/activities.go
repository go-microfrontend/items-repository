package processes

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/pkg/errors"

	repo "github.com/go-microfrontend/items-repository/internal/repository"
)

type Repo interface {
	CreateItem(ctx context.Context, arg repo.CreateItemParams) (pgtype.UUID, error)
	GetItemByID(ctx context.Context, id pgtype.UUID) (repo.Item, error)
	GetItems(ctx context.Context, arg repo.GetItemsParams) ([]repo.Item, error)
}

type Activities struct {
	repo Repo
}

func New(repo Repo) *Activities {
	return &Activities{repo: repo}
}

func (a *Activities) CreateItem(ctx context.Context, arg repo.CreateItemParams) (string, error) {
	id, err := a.repo.CreateItem(ctx, arg)
	if err != nil {
		return "", errors.Wrap(err, "creating item")
	}

	return id.String(), nil
}

func (a *Activities) GetItemByID(ctx context.Context, id string) (*repo.Item, error) {
	var pgid pgtype.UUID
	err := pgid.Scan(id)
	if err != nil {
		return nil, errors.Wrap(err, "getting item by id")
	}

	item, err := a.repo.GetItemByID(ctx, pgid)
	if err != nil {
		return nil, errors.Wrap(err, "gettings item by id")
	}

	return &item, nil
}

func (a *Activities) GetItems(ctx context.Context, arg repo.GetItemsParams) ([]repo.Item, error) {
	items, err := a.repo.GetItems(ctx, arg)
	if err != nil {
		return nil, errors.Wrap(err, "getting items")
	}

	return items, nil
}
