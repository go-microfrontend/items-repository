package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"

	repo "github.com/go-microfrontend/items-repository/internal/repository"
)

type Cache struct {
	client *redis.Client
	ttl    time.Duration
}

func NewCache(client *redis.Client, ttl time.Duration) *Cache {
	return &Cache{
		client: client,
		ttl:    ttl,
	}
}

func (r *Cache) GetCategories(ctx context.Context) ([]repo.Category, error) {
	key := "categories"
	val, err := r.client.Get(ctx, key).Result()
	if err == nil {
		var categories []repo.Category
		if err := json.Unmarshal([]byte(val), &categories); err == nil {
			return categories, nil
		}
	}

	return nil, redis.Nil
}

func (r *Cache) GetProductsByCategory(
	ctx context.Context,
	category *string,
) ([]repo.Product, error) {
	if category == nil {
		return nil, redis.Nil
	}

	key := "products:" + *category
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var products []repo.Product
	if err := json.Unmarshal([]byte(val), &products); err != nil {
		return nil, err
	}
	return products, nil
}

func (r *Cache) GetProductByID(
	ctx context.Context,
	productID uuid.UUID,
) (repo.Product, error) {
	key := fmt.Sprintf("product:%s", productID)
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return repo.Product{}, redis.Nil
	}

	var product repo.Product
	if err := json.Unmarshal([]byte(val), &product); err != nil {
		return repo.Product{}, redis.Nil
	}
	return product, nil
}

func (r *Cache) GetProductCharacteristicByID(
	ctx context.Context,
	productID uuid.UUID,
) (repo.ProductCharacteristic, error) {
	key := fmt.Sprintf("product_char:%s", productID)
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return repo.ProductCharacteristic{}, redis.Nil
	}

	var char repo.ProductCharacteristic
	if err := json.Unmarshal([]byte(val), &char); err != nil {
		return repo.ProductCharacteristic{}, redis.Nil
	}
	return char, nil
}
