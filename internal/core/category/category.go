// Package beer provides an example of a core business API. Right now these
// calls are just wrapping the data/store layer. But at some point you will
// want to audit or something that isn't specific to the data/store layer.
package category

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/chagasVinicius/apollo/internal/sys/database"
	"github.com/chagasVinicius/apollo/internal/sys/validate"
	"github.com/google/uuid"
)

// Set of error variables for CRUD operations.
var (
	ErrNotFound  = errors.New("category not found")
	ErrInvalidID = errors.New("ID is not in its proper form")
)

// Storer .
type Storer interface {
	AddCategory(ctx context.Context, category Category) error
	QueryCategories(ctx context.Context, page int, size int) ([]Category, error)
	QueryCategoryByID(ctx context.Context, categoryID string) (Category, error)
}

// Core manages the set of APIs for beer access.
type Core struct {
	store Storer
}

// NewCore constructs a core for product api access.
func NewCore(store Storer) Core {
	return Core{
		store: store,
	}
}

// =========================================================================
// Category Support

// Create adds an category to the database. Its return the created Category
// with fields populated.
func (c Core) Create(ctx context.Context, nc NewCategory) (Category, error) {
	if err := validate.Check(nc); err != nil {
		return Category{}, fmt.Errorf("validating data: %w", err)
	}

	category := Category{
		ID:        uuid.New().String(),
		Name:      nc.Name,
		ShortDesc: nc.ShortDesc,
		CreatedAt: time.Now(),
	}

	if err := c.store.AddCategory(ctx, category); err != nil {
		return Category{}, fmt.Errorf("addCategory: %w", err)
	}

	return category, nil
}

// QueryByID gets the specified category from the database.
func (c Core) QueryByID(ctx context.Context, id string) (Category, error) {
	if err := validate.CheckID(id); err != nil {
		return Category{}, ErrInvalidID
	}

	category, err := c.store.QueryCategoryByID(ctx, id)
	if err != nil {
		if database.IsNoRowError(err) {
			return Category{}, ErrNotFound
		}
		return Category{}, fmt.Errorf("queryCategoryByID: %w", err)
	}

	return category, nil
}

// Query gets all categories from the database.
func (c Core) Query(ctx context.Context, page, pageSize int) ([]Category, error) {
	categories, err := c.store.QueryCategories(ctx, page, pageSize)
	if err != nil {
		return nil, fmt.Errorf("queryCategories: %w", err)
	}

	return categories, nil
}
