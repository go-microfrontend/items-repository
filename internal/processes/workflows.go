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

var Workflows = []any{
	GetCategories,
	GetProductsByCategory,
	GetProductByID,
	GetProductCharacteristicByID,
}

func GetCategories(ctx workflow.Context) ([]repo.Category, error) {
	ctx = workflow.WithActivityOptions(ctx, itemsActivityOptions)

	var categories []repo.Category
	err := workflow.ExecuteActivity(ctx, "GetCategories").Get(ctx, &categories)
	if err != nil {
		return nil, errors.Wrap(err, "executing GetCategories activity")
	}

	return categories, nil
}

func GetProductsByCategory(ctx workflow.Context, category string) ([]repo.Product, error) {
	ctx = workflow.WithActivityOptions(ctx, itemsActivityOptions)

	var products []repo.Product
	err := workflow.ExecuteActivity(ctx, "GetProductsByCategory", category).Get(ctx, &products)
	if err != nil {
		return nil, errors.Wrap(err, "executing GetCategories activity")
	}

	return products, nil
}

func GetProductByID(ctx workflow.Context, id string) (*repo.Product, error) {
	ctx = workflow.WithActivityOptions(ctx, itemsActivityOptions)

	var product repo.Product
	err := workflow.ExecuteActivity(ctx, "GetProductByID", id).Get(ctx, &product)
	if err != nil {
		return nil, errors.Wrap(err, "executing GetProductByID activity")
	}

	return &product, nil
}

func GetProductCharacteristicByID(
	ctx workflow.Context,
	id string,
) (*repo.ProductCharacteristic, error) {
	ctx = workflow.WithActivityOptions(ctx, itemsActivityOptions)

	var characteristic repo.ProductCharacteristic
	err := workflow.ExecuteActivity(ctx, "GetProductCharacteristicByID", id).
		Get(ctx, &characteristic)
	if err != nil {
		return nil, errors.Wrap(err, "executing GetProductCharacteristicByID activity")
	}

	return &characteristic, nil
}
