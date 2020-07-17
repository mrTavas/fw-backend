package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"

	//local
	"github.com/labstack/echo"
	db "github.com/mrTavas/fw-backend/dbconn"
	"github.com/mrTavas/fw-backend/models"
)

// UploadWorkerImage  - upload all files in /var/www/html/uploads/workersImages
func UploadWorkerImage(c echo.Context) error {

	// Read form fields
	workerUUID := c.FormValue("workerUUID")

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	if file.Filename[len(file.Filename)-3:len(file.Filename)] != "png" && file.Filename[len(file.Filename)-3:len(file.Filename)] != "jpg" && file.Filename[len(file.Filename)-4:len(file.Filename)] != "jpeg" {
		return c.HTML(http.StatusOK, fmt.Sprintf("Error. (.png, .jpg, .jpeg formats only)"))

	}

	// Clear folder
	cmd := exec.Command("rm", "-rf", "/var/www/html/uploads/workersImages"+workerUUID)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	err = os.MkdirAll("/var/www/html/uploads/workersImages/"+workerUUID, 0777)
	if err != nil {
		return err
	}

	dst, err := os.Create("/var/www/html/uploads/workersImages/" + workerUUID + "/" + file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	var worker models.Workers

	_, err = db.Conn.Query(&worker, "SELECT * FROM workers WHERE uuid = ?", workerUUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	worker.ImageLink = "http://fwqqq-backend.ddns.net:8001/uploads/workersImages/" + workerUUID + "/" + file.Filename

	err = db.Conn.Update(&worker)
	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully</p>", file.Filename))
}
