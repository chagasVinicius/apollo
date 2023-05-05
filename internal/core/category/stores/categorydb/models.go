package categorydb

import (
	"time"

	"github.com/chagasVinicius/apollo/internal/core/category"
	"github.com/uptrace/bun"
)

// dbCategory represents an individual category.
type dbCategory struct {
	bun.BaseModel `bun:"table:categories,alias:c"`

	ID        string    `bun:"id,pk"`
	Name      string    `bun:"name"`
	ShortDesc string    `bun:"short_desc"`
	CreatedAt time.Time `bun:"created_at"`
}

// =========================================================

func toDBCategory(c category.Category) dbCategory {
	return dbCategory{
		ID:        c.ID,
		Name:      c.Name,
		ShortDesc: c.ShortDesc,
		CreatedAt: c.CreatedAt,
	}
}

func toCategory(dbC dbCategory) category.Category {
	return category.Category{
		ID:        dbC.ID,
		Name:      dbC.Name,
		ShortDesc: dbC.ShortDesc,
		CreatedAt: dbC.CreatedAt,
	}
}

func toCategories(list []dbCategory) []category.Category {
	categories := make([]category.Category, len(list))
	for i, c := range list {
		categories[i] = toCategory(c)
	}
	return categories
}
