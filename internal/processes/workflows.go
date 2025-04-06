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

var Workflows = []any{CreateItemWF, GetItemByIDWF, GetItemsWF}

func CreateItemWF(ctx workflow.Context, arg repo.CreateItemParams) error {
	ctx = workflow.WithActivityOptions(ctx, itemsActivityOptions)

	err := workflow.ExecuteActivity(ctx, "CreateItem", arg).Get(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "executing CreateItem activity")
	}

	return nil
}

func GetItemByIDWF(ctx workflow.Context, id string) (*repo.Item, error) {
	ctx = workflow.WithActivityOptions(ctx, itemsActivityOptions)

	var item repo.Item
	err := workflow.ExecuteActivity(ctx, "executing GetItemByID activity", id).Get(ctx, &item)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func GetItemsWF(ctx workflow.Context, arg repo.GetItemsParams) ([]repo.Item, error) {
	ctx = workflow.WithActivityOptions(ctx, itemsActivityOptions)

	var items []repo.Item
	err := workflow.ExecuteActivity(ctx, "executing GetItems activity", arg).Get(ctx, &items)
	if err != nil {
		return nil, err
	}

	return items, nil
}
