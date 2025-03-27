package server

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/go-microfrontend/items-repository/internal/repository"
	desc "github.com/go-microfrontend/items-repository/pkg/items"
)

func (s Server) CreateItem(
	ctx context.Context,
	req *desc.CreateItemRequest,
) (*desc.CreateItemResponse, error) {
	id, err := s.repo.CreateItem(ctx, repository.CreateItemParams{
		Name:          req.GetName(),
		Description:   req.GetDescription(),
		Type:          req.GetType(),
		WeightInGrams: req.GetWeightInGrams(),
	})
	if err != nil {
		return nil, status.Error(
			codes.Internal,
			fmt.Errorf("[api::CreateItem] database: %w", err).Error(),
		)
	}

	return &desc.CreateItemResponse{
		Id: id.String(),
	}, nil
}

func (s Server) GetItemByID(
	ctx context.Context,
	req *desc.GetItemByIDRequest,
) (*desc.Item, error) {
	var id pgtype.UUID
	err := id.Scan(req.Id)
	if err != nil {
		return nil, status.Error(
			codes.InvalidArgument,
			fmt.Errorf("[api::GetItemByID] uuid: %w", err).Error(),
		)
	}

	item, err := s.repo.GetItemByID(ctx, id)
	if err != nil {
		return nil, status.Error(
			codes.Internal,
			fmt.Errorf("[api::CreateItem] database: %w", err).Error(),
		)
	}

	return &desc.Item{
		Id:            item.ID.String(),
		Name:          item.Name,
		Description:   item.Description,
		Type:          item.Type,
		WeightInGrams: item.WeightInGrams,
	}, nil
}

func (s Server) GetItems(
	ctx context.Context,
	req *desc.GetItemsRequest,
) (*desc.GetItemsResponse, error) {
	items, err := s.repo.GetItems(ctx, repository.GetItemsParams{
		Limit:  req.GetLimit(),
		Offset: req.GetOffset(),
	})
	if err != nil {
		return nil, status.Error(
			codes.Internal,
			fmt.Errorf("[api::CreateItem] database: %w", err).Error(),
		)
	}

	resp := &desc.GetItemsResponse{
		Items: make([]*desc.Item, 0, len(items)),
	}
	for _, item := range items {
		resp.Items = append(resp.Items, &desc.Item{
			Id:            item.ID.String(),
			Name:          item.Name,
			Description:   item.Description,
			Type:          item.Type,
			WeightInGrams: item.WeightInGrams,
		})
	}

	return resp, nil
}
