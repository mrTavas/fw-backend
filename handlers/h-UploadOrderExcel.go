package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"

	//local
	"github.com/labstack/echo"
)

// UploadOrderExcel  - upload order excel list /var/www/html/uploads/?
func UploadOrderExcel(c echo.Context) error {

	// Read form fields
	name := c.FormValue("name")

	println(name)

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	err = os.MkdirAll("/home/tavas/"+name, 0777)
	if err != nil {
		return err
	}

	dst, err := os.Create("/home/tavas/" + name + "/" + file.Filename)
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
