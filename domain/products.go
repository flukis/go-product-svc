package domain

import (
	"context"
	"time"
)

type Products struct {
	ID          string    `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
	Name        string    `db:"product_name" json:"name" validate:"required,lte=255"`
	Description string    `db:"product_desc" json:"description" validate:"required,lte=255"`
	ImageSrc    string    `json:"product_img_src"`
}

type ProductUsecase interface {
	Fetch(ctx context.Context, cursor string, num int64) ([]Products, string, error)
	GetByID(ctx context.Context, id string) (Products, error)
	GetByName(ctx context.Context, name string) (Products, error)
	Store(ctx context.Context, p *Products) error
	Update(ctx context.Context, p *Products) error
	Delete(ctx context.Context, id string) error
}

type ProductRepository interface {
	Fetch(ctx context.Context, cursor string, num int64) (res []Products, nextCursor string, err error)
	GetByID(ctx context.Context, id string) (res Products, err error)
	GetByName(ctx context.Context, name string) (res Products, err error)
	Store(ctx context.Context, p *Products) (err error)
	Update(ctx context.Context, p *Products) (err error)
	Delete(ctx context.Context, id string) (err error)
}
