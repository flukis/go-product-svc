package usecases

import (
	"context"
	"time"

	"github.com/fahmilukis/go-product-svc/domain"
)

type productUsecase struct {
	productRepository domain.ProductRepository
	ctxTimeout        time.Duration
}

func NewProductUsecase(p domain.ProductRepository, to time.Duration) domain.ProductUsecase {
	return &productUsecase{
		productRepository: p,
		ctxTimeout:        to,
	}
}

func (p *productUsecase) Fetch(c context.Context, cursor string, num int64) (res []domain.Products, nextCursor string, err error) {
	if num < 10 {
		num = 10
	}

	ctx, cancel := context.WithTimeout(c, p.ctxTimeout)
	defer cancel()

	res, nextCursor, err = p.productRepository.Fetch(ctx, cursor, num)
	if err != nil {
		return nil, "", err
	}

	return
}

func (p *productUsecase) GetByID(c context.Context, id string) (res domain.Products, err error) {
	ctx, cancel := context.WithTimeout(c, p.ctxTimeout)
	defer cancel()

	res, err = p.productRepository.GetByID(ctx, id)
	if err != nil {
		return domain.Products{}, err
	}

	return
}

func (p *productUsecase) GetByName(c context.Context, name string) (res domain.Products, err error) {
	ctx, cancel := context.WithTimeout(c, p.ctxTimeout)
	defer cancel()

	res, err = p.productRepository.GetByName(ctx, name)
	if err != nil {
		return domain.Products{}, err
	}

	return
}

func (p *productUsecase) Store(c context.Context, a *domain.Products) (err error) {
	ctx, cancel := context.WithTimeout(c, p.ctxTimeout)
	defer cancel()

	now := time.Now()
	a.CreatedAt = now
	a.UpdatedAt = now

	return p.productRepository.Store(ctx, a)
}

func (p *productUsecase) Update(c context.Context, a *domain.Products) (err error) {
	ctx, cancel := context.WithTimeout(c, p.ctxTimeout)
	defer cancel()

	a.UpdatedAt = time.Now()

	return p.productRepository.Store(ctx, a)
}

func (p *productUsecase) Delete(c context.Context, id string) (err error) {
	ctx, cancel := context.WithTimeout(c, p.ctxTimeout)
	defer cancel()

	return p.productRepository.Delete(ctx, id)
}
