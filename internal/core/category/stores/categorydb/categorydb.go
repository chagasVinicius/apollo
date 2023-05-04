package categorydb

import (
	"context"
	"fmt"

	"github.com/chagasVinicius/apollo/internal/core/category"
	"github.com/uptrace/bun"
	"go.uber.org/zap"
)

// Store manages the set of APIs for category access.
type Store struct {
	log *zap.SugaredLogger
	db  *bun.DB
}

// NewStore constructs a data for api access.
func NewStore(log *zap.SugaredLogger, db *bun.DB) Store {
	return Store{
		log: log,
		db:  db,
	}
}

// AddCategory adds a new category to the database.
func (s Store) AddCategory(ctx context.Context, b category.Category) error {
	dbCategory := toDBCategory(b)

	if _, err := s.db.NewInsert().Model(&dbCategory).Exec(ctx); err != nil {
		return fmt.Errorf("adding category: %w", err)
	}

	return nil
}

// QueryCategoryByID retrieves a category by its id.
func (s Store) QueryCategoryByID(ctx context.Context, categoryID string) (category.Category, error) {
	var b dbCategory

	query := s.db.NewSelect().
		Model(&b).
		Where("id = ?", categoryID)

	if err := query.Scan(ctx); err != nil {
		return category.Category{}, fmt.Errorf("querying category by [id=%s]: %w", categoryID, err)
	}

	return toCategory(b), nil
}

// QueryCategories retrieves a list of existing categories.
func (s Store) QueryCategories(ctx context.Context, page, size int) ([]category.Category, error) {
	var categories []dbCategory

	query := s.db.NewSelect().
		Model(&categories).
		Limit(size).
		Offset(size * (page - 1))

	if err := query.Scan(ctx); err != nil {
		return nil, fmt.Errorf("querying category: %w", err)
	}

	return toCategories(categories), nil
}
