package usecases

import (
	"context"
	"time"

	"github.com/fahmilukis/go-product-svc/domain"
	pkg "github.com/fahmilukis/go-product-svc/pkg/utils"
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

func (p *productUsecase) Fetch(c context.Context, pg pkg.Pagination) (res []domain.Products, nextPkg pkg.Pagination, err error) {
	ctx, cancel := context.WithTimeout(c, p.ctxTimeout)
	defer cancel()

	res, nextPkg, err = p.productRepository.Fetch(ctx, pg)
	if err != nil {
		return nil, pkg.Pagination{}, err
	}
	if len(res) == 0 {
		return res, pkg.Pagination{}, domain.ErrNotFound
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
