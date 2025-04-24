package processes

import (
	"time"

	"github.com/pkg/errors"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"

	repo "github.com/go-microfrontend/items-repository/internal/repository"
)

var itemsActivityOptions = workflow.ActivityOptions{
	StartToCloseTimeout: time.Minute,
	RetryPolicy: &temporal.RetryPolicy{
		InitialInterval:    time.Second,
		BackoffCoefficient: 2.0,
		MaximumInterval:    10 * time.Second,
		MaximumAttempts:    5,
	},
}

var Workflows = []any{CreateItem, GetItemByID, GetItems, GetItemsByType}

func CreateItem(ctx workflow.Context, arg repo.CreateItemParams) (string, error) {
	ctx = workflow.WithActivityOptions(ctx, itemsActivityOptions)

	var id string
	err := workflow.ExecuteActivity(ctx, "CreateItem", arg).Get(ctx, &id)
	if err != nil {
		return "", errors.Wrap(err, "executing CreateItem activity")
	}

	return id, nil
}

func GetItemByID(ctx workflow.Context, id string) (*repo.Item, error) {
	ctx = workflow.WithActivityOptions(ctx, itemsActivityOptions)

	var item repo.Item
	err := workflow.ExecuteActivity(ctx, "GetItemByID", id).Get(ctx, &item)
	if err != nil {
		return nil, errors.Wrap(err, "executing GetItemByID activity")
	}

	return &item, nil
}

func GetItems(ctx workflow.Context, arg repo.GetItemsParams) ([]repo.Item, error) {
	ctx = workflow.WithActivityOptions(ctx, itemsActivityOptions)

	var items []repo.Item
	err := workflow.ExecuteActivity(ctx, "GetItems", arg).Get(ctx, &items)
	if err != nil {
		return nil, errors.Wrap(err, "executing GetItems activity")
	}

	return items, nil
}

func GetItemsByType(ctx workflow.Context, arg repo.GetItemsByTypeParams) ([]repo.Item, error) {
	ctx = workflow.WithActivityOptions(ctx, itemsActivityOptions)

	var items []repo.Item
	err := workflow.ExecuteActivity(ctx, "GetItemsByType", arg).Get(ctx, &items)
	if err != nil {
		return nil, errors.Wrap(err, "executing GetItemsByType activity")
	}

	return items, nil
}
