// Package categorygrp maintains the group of handlers for category access.
package categorygrp

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/chagasVinicius/apollo/internal/core/category"
	v1Web "github.com/chagasVinicius/apollo/internal/web/v1"
	"github.com/chagasVinicius/apollo/kit/web"
)

const (
	defaultPage = 1
	defaultSize = 10
)

// Handlers manages the set of category endpoints.
type Handlers struct {
	Category category.Core
}

// Create adds a new category to the system.
func (h Handlers) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var nc category.NewCategory
	if err := web.Decode(r, &nc); err != nil {
		return fmt.Errorf("unable to decode payload: %w", err)
	}

	c, err := h.Category.Create(ctx, nc)
	if err != nil {
		return fmt.Errorf("creating new category, nb[%+v]: %w", nc, err)
	}

	return web.Respond(ctx, w, c, http.StatusCreated)
}

// QueryByID returns a category by its ID.
func (h Handlers) QueryByID(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	id := web.Param(r, "id")
	c, err := h.Category.QueryByID(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, category.ErrInvalidID):
			return v1Web.NewRequestError(err, http.StatusBadRequest)
		case errors.Is(err, category.ErrNotFound):
			return v1Web.NewRequestError(err, http.StatusNotFound)
		default:
			return fmt.Errorf("ID[%s]: %w", id, err)
		}
	}

	return web.Respond(ctx, w, c, http.StatusOK)
}

func (h Handlers) Query(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	page := web.Query(r, "page", defaultPage)
	pageNumber, err := strconv.Atoi(page)
	if err != nil {
		return v1Web.NewRequestError(fmt.Errorf("invalid page format, page[%s]", page), http.StatusBadRequest)
	}

	size := web.Query(r, "size", defaultSize)
	sizeNumber, err := strconv.Atoi(size)
	if err != nil {
		return v1Web.NewRequestError(fmt.Errorf("invalid rows format, size[%s]", size), http.StatusBadRequest)
	}

	list, err := h.Category.Query(ctx, pageNumber, sizeNumber)
	if err != nil {
		return fmt.Errorf("querying categories: %w", err)
	}

	if len(list) == 0 {
		return web.Respond(ctx, w, nil, http.StatusNoContent)
	}

	return web.Respond(ctx, w, list, http.StatusOK)
}
