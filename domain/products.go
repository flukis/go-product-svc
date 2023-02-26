package domain

import (
	"context"
	"time"
)

type Products struct {
	ID          int64     `json:"id"`
	Name        string    `json:"product_name"`
	Description string    `json:"product_description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	ImageSrc    string    `json:"product_img_src"`
}

type ProductUsecase interface {
	Fetch(ctx context.Context, cursor string, num int64) ([]Products, error)
	GetByID(ctx context.Context, id int64) (Products, error)
	GetByName(ctx context.Context, name string) (Products, error)
	Store(ctx context.Context, p *Products) error
	Update(ctx context.Context, p *Products) error
	Delete(ctx context.Context, id int64) error
}

type ProductRepository interface {
	Fetch(ctx context.Context, cursor string, num int64) (res []Products, nextCursor string, err error)
	GetByID(ctx context.Context, id int64) (res Products, err error)
	GetByName(ctx context.Context, name string) (res Products, err error)
	Store(ctx context.Context, p *Products) (err error)
	Update(ctx context.Context, p *Products) (err error)
	Delete(ctx context.Context, id int64) (err error)
}
