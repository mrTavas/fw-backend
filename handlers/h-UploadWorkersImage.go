package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"

	//local
	"github.com/labstack/echo"
)

// UploadWorkerImage  - upload all files in /var/www/html/uploads/workersImages
func UploadWorkerImage(c echo.Context) error {

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	if file.Filename[len(file.Filename)-3:len(file.Filename)] != "png" && file.Filename[len(file.Filename)-3:len(file.Filename)] != "jpg" && file.Filename[len(file.Filename)-4:len(file.Filename)] != "jpeg" {
		return c.HTML(http.StatusOK, fmt.Sprintf("Error. (.png, .jpg, .jpeg formats only)"))

	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create("/var/www/html/uploads/workersImages/" + file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully</p>", file.Filename))
}
