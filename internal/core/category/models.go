package category

import (
	"time"
)

type Category struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	ShortDesc string    `json:"short_desc"`
	CreatedAt time.Time `json:"created_at"`
}

type NewCategory struct {
	Name      string `json:"name" validate:"required"`
	ShortDesc string `json:"short_desc" validate:"required"`
}
