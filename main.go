package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/fahmilukis/go-product-svc/files"
	pkg "github.com/fahmilukis/go-product-svc/pkg/utils"
	handler "github.com/fahmilukis/go-product-svc/products/handler/http"
	"github.com/fahmilukis/go-product-svc/products/repositories"
	"github.com/fahmilukis/go-product-svc/products/usecases"
	"github.com/gofiber/fiber/v2"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgresql://postgres:wap12345@localhost:5432/postgres?sslmode=disable"
	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	app := fiber.New()
	files.NewUploadImageRoutes(app)

	productRepo := repositories.NewProductDBRepository(dbConn)
	productUsecase := usecases.NewProductUsecase(productRepo, 10*time.Second)

	handler.ProductRoute(app, productUsecase)

	pkg.StartServer(app)
}
