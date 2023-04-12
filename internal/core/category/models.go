package category

import (
	"fmt"
)

type Category struct {
	ID string `json:"id"`
	Name string `json:"name"`
	ShortDesc string `json:"short_desc"`
	CreatedAt string `json:"created_at"`
}

type NewCategory struct {
	Name string `json:"name"`
	ShortDesc string `json:"short_desc"`

}
