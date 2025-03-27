package server

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"

	"github.com/go-microfrontend/items-repository/internal/repository"
	desc "github.com/go-microfrontend/items-repository/pkg/items"
)

type Repo interface {
	CreateItem(ctx context.Context, arg repository.CreateItemParams) (pgtype.UUID, error)
	GetItemByID(ctx context.Context, id pgtype.UUID) (repository.Item, error)
	GetItems(ctx context.Context, arg repository.GetItemsParams) ([]repository.Item, error)
}

type Server struct {
	desc.UnimplementedItemServiceServer
	repo Repo
}

func New(repo Repo) *Server {
	return &Server{repo: repo}
}
