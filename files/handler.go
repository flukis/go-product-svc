package files

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

const MAX_UPLOAD_SIZE = 1024 * 1024 * 5

func NewUploadImageRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	route.Post("/uploader/image", UploadImage)
	route.Get("/download/:id", GetImage)
}

func GetImage(c *fiber.Ctx) error {
	fileName := c.Params("id")
	filePath := filepath.Join(os.TempDir(), "image_server", fileName)
	return c.Download(filePath, fileName)
}

func UploadImage(c *fiber.Ctx) error {
	file, err := c.FormFile("document")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status": false,
			"msg":    err,
			"data":   "",
		})
	}

	generateFilename := tempFileName(file.Filename)
	filePath := filepath.Join(os.TempDir(), "image_server", generateFilename)

	if err = c.SaveFile(file, filePath); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status": false,
			"msg":    err,
			"data":   "",
		})
	}

	return c.JSON(fiber.Map{
		"status": true,
		"msg":    "success upload image",
		"data":   fmt.Sprintf("%s/api/v1/download/%s", c.BaseURL(), generateFilename),
	})
}

func tempFileName(fileName string) string {
	randBytes := make([]byte, 16)

	rand.Read(randBytes)
	suffix := fileName[len(fileName)-len(filepath.Ext(fileName)):]
	prefix := fileName[:len(fileName)-len(filepath.Ext(fileName))]
	return fmt.Sprint(prefix + "-" + hex.EncodeToString(randBytes) + suffix)
}
