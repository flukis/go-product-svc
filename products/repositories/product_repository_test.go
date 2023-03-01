package repositories_test

import (
	"context"
	"testing"
	"time"

	"github.com/fahmilukis/go-product-svc/domain"
	pkg "github.com/fahmilukis/go-product-svc/pkg/utils"
	"github.com/fahmilukis/go-product-svc/products/repositories"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestFetchRepositoryProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	now := time.Now()

	mockProducts := []domain.Products{
		{
			ID:          "1",
			Name:        "product 1",
			Description: "description 1",
			CreatedAt:   now,
			UpdatedAt:   now,
			ImageSrc:    "img_src_1",
		},
		{
			ID:          "2",
			Name:        "product 2",
			Description: "description 2",
			CreatedAt:   now,
			UpdatedAt:   now,
			ImageSrc:    "img_src_2",
		},
	}

	rows := sqlmock.NewRows([]string{"id", "product_name", "product_desc", "created_at", "updated_at", "product_img_src"}).
		AddRow(mockProducts[0].ID, mockProducts[0].Name, mockProducts[0].Description, mockProducts[0].CreatedAt, mockProducts[0].UpdatedAt, mockProducts[0].ImageSrc).
		AddRow(mockProducts[1].ID, mockProducts[1].Name, mockProducts[1].Description, mockProducts[1].CreatedAt, mockProducts[1].UpdatedAt, mockProducts[1].ImageSrc)

	query := `SELECT id,product_name,product_desc,created_at,updated_at,product_img_src
	FROM products WHERE created_at > \$1 ORDER BY created_at LIMIT \$2`

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := repositories.NewProductDBRepository(db)
	cursor := pkg.EncodeCursor(mockProducts[1].CreatedAt)
	num := int64(2)
	list, nextCursor, err := a.Fetch(context.TODO(), cursor, num)
	assert.NotEmpty(t, nextCursor)
	assert.NoError(t, err)
	assert.Len(t, list, 2)
}

func TestInsertRepositoryProduct(t *testing.T) {
	now := time.Now()
	ar := &domain.Products{
		Name:        "product test",
		Description: "description test",
		CreatedAt:   now,
		UpdatedAt:   now,
		ImageSrc:    "img_url_3",
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := `INSERT INTO products \(product_name,product_desc,created_at,updated_at,product_img_src\) VALUES \(\$1, \$2, \$3, \$4, \$5\)`
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(ar.Name, ar.Description, ar.CreatedAt, ar.UpdatedAt, ar.ImageSrc).WillReturnResult(sqlmock.NewResult(12, 1))

	a := repositories.NewProductDBRepository(db)

	err = a.Store(context.TODO(), ar)
	assert.NoError(t, err)
	assert.Equal(t, "12", ar.ID)
}

func TestUpdateRepositoryProduct(t *testing.T) {
	now := time.Now()
	ar := &domain.Products{
		ID:          "12",
		Name:        "product test",
		Description: "description test",
		CreatedAt:   now,
		UpdatedAt:   now,
		ImageSrc:    "img_url_3",
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := `UPDATE products SET product_name=\$1 , product_desc=\$2 , created_at=\$3 , updated_at=\$4 , product_img_src=\$5  WHERE id \$6`
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(ar.Name, ar.Description, ar.CreatedAt, ar.UpdatedAt, ar.ImageSrc, ar.ID).WillReturnResult(sqlmock.NewResult(12, 1))

	a := repositories.NewProductDBRepository(db)

	err = a.Update(context.TODO(), ar)
	assert.NoError(t, err)
}

func TestDeleteRepositoryProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := `DELETE FROM products WHERE id = \$1`

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs("12").WillReturnResult(sqlmock.NewResult(12, 1))

	a := repositories.NewProductDBRepository(db)

	num := "12"
	err = a.Delete(context.TODO(), num)
	assert.NoError(t, err)
}

func TestGetByIdRepositoryProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	now := time.Now().Add(1)

	rows := sqlmock.NewRows([]string{
		"id", "product_name", "product_desc", "created_at", "updated_at", "product_img_src",
	}).AddRow(
		"3", "product 1", "desc 1", now, now, "img_src",
	)

	mockData := domain.Products{
		ID:          "3",
		Name:        "product 1",
		Description: "desc 1",
		CreatedAt:   now,
		UpdatedAt:   now,
		ImageSrc:    "img_src",
	}

	query := `SELECT id,product_name,product_desc,created_at,updated_at,product_img_src from products WHERE id=\$1`

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := repositories.NewProductDBRepository(db)

	num := "3"
	aProduct, err := a.GetByID(context.TODO(), num)
	assert.NoError(t, err)
	assert.NotNil(t, aProduct)
	assert.Equal(t, mockData, aProduct)
}
