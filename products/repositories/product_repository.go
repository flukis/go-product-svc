package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"math"

	"github.com/fahmilukis/go-product-svc/domain"
	pkg "github.com/fahmilukis/go-product-svc/pkg/utils"
	"github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

type productDBRepositories struct {
	Conn *sql.DB
}

func NewProductDBRepository(conn *sql.DB) *productDBRepositories {
	return &productDBRepositories{Conn: conn}
}

// fetch to DB
func (p *productDBRepositories) fetch(ctx context.Context, query string, args ...interface{}) (res []domain.Products, err error) {
	rows, err := p.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(err)
		}
	}()

	res = make([]domain.Products, 0)
	for rows.Next() {
		prd := domain.Products{}
		err = rows.Scan(
			&prd.ID,
			&prd.Name,
			&prd.Description,
			&prd.CreatedAt,
			&prd.UpdatedAt,
			&prd.ImageSrc,
		)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		res = append(res, prd)
	}

	return res, nil
}

func (p *productDBRepositories) Fetch(ctx context.Context, pagination pkg.Pagination) (res []domain.Products, nextPagination pkg.Pagination, err error) {
	query := `SELECT id,product_name,product_desc,created_at,updated_at,product_img_src
	FROM products ORDER BY created_at ASC LIMIT $1 OFFSET $2`

	res, err = p.fetch(ctx, query, pagination.Limit, pagination.GetOffset())
	if err != nil {
		return nil, pkg.Pagination{}, err
	}

	var total int64
	p.Conn.QueryRow(`SELECT reltuples AS estimate FROM pg_class where relname = 'products'`).Scan(&total)
	nextPagination.TotalRows = total

	totalPages := int(math.Ceil(float64(total) / float64(pagination.Limit)))
	nextPagination.TotalPages = totalPages

	nextPagination.Limit = pagination.Limit
	nextPagination.Page = pagination.Page

	return
}

func (p *productDBRepositories) GetByID(ctx context.Context, id string) (res domain.Products, err error) {
	query := `SELECT id,product_name,product_desc,created_at,updated_at,product_img_src from products WHERE id=$1`
	list, err := p.fetch(ctx, query, id)
	if err != nil {
		return domain.Products{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}

func (p *productDBRepositories) GetByName(ctx context.Context, name string) (res domain.Products, err error) {
	query := `SELECT id,product_name,product_desc,created_at,updated_at,product_img_src from products WHERE product_name=?`
	list, err := p.fetch(ctx, query, name)
	if err != nil {
		return domain.Products{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}

// insert a record
func (p *productDBRepositories) Store(ctx context.Context, prd *domain.Products) (err error) {
	query := `INSERT INTO products (product_name,product_desc,created_at,updated_at,product_img_src) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	stmt, err := p.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	row := stmt.QueryRowContext(
		ctx,
		prd.Name,
		prd.Description,
		prd.CreatedAt,
		prd.UpdatedAt,
		prd.ImageSrc,
	)
	var id string
	if err = row.Scan(&id); err != nil {
		return
	}

	prd.ID = id
	return
}

func (p *productDBRepositories) Update(ctx context.Context, prd *domain.Products) (err error) {
	query := `UPDATE products SET product_name=$1 , product_desc=$2 , created_at=$3 , updated_at=$4 , product_img_src=$5  WHERE id $6`

	stmt, err := p.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, prd.Name, prd.Description, prd.CreatedAt, prd.UpdatedAt, prd.ImageSrc, prd.ID)
	if err != nil {
		return
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return
	}
	if affect != 1 {
		err = fmt.Errorf("weird  Behavior. Total Affected: %d", affect)
		return
	}

	return
}

func (p *productDBRepositories) Delete(ctx context.Context, id string) (err error) {
	query := "DELETE FROM products WHERE id = $1"

	stmt, err := p.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return
	}

	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return
	}

	if rowsAfected != 1 {
		err = fmt.Errorf("weird  Behavior. Total Affected: %d", rowsAfected)
		return
	}

	return
}
