package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"

	//local
	"github.com/labstack/echo"
	db "github.com/mrTavas/fw-backend/dbconn"
	"github.com/mrTavas/fw-backend/models"
)

// UploadOrderExcel  - upload order excel list /var/www/html/uploads/?
func UploadOrderExcel(c echo.Context) error {

	var order models.Orders
	// Read form fields
	orderID := c.FormValue("orderID")

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	println(file.Filename[len(file.Filename)-4 : len(file.Filename)])
	if file.Filename[len(file.Filename)-4:len(file.Filename)] != "xlsx" && file.Filename[len(file.Filename)-4:len(file.Filename)] != "xlsm" && file.Filename[len(file.Filename)-3:len(file.Filename)] != "xls" {
		return c.HTML(http.StatusOK, fmt.Sprintf("Not Excel file. (.xlsx, .xlsm, .xls formats only.)"))

	}

	// Clear folder
	cmd := exec.Command("rm", "-rf", "/var/www/html/uploads/orders/"+orderID)
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
	err = os.MkdirAll("/var/www/html/uploads/orders/"+orderID, 0777)
	if err != nil {
		return err
	}

	dst, err := os.Create("/var/www/html/uploads/orders/" + orderID + "/" + file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	// Convert to pdf
	cmd = exec.Command("unoconv", "-f", "pdf", "/var/www/html/uploads/orders/"+orderID+"/"+file.Filename)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()

	lenSymbols := 4
	if file.Filename[len(file.Filename)-3:len(file.Filename)] == "xls" {
		lenSymbols = 3
	}

	orderIDint, err := strconv.Atoi(orderID)
	err = db.Conn.Insert(&models.OrdersFilesLinks{
		OrderID: orderIDint,
		Excel:   "http://fwqqq-backend.ddns.net:8001/uploads/orders/" + orderID + "/" + file.Filename,
		Pdf:     "http://fwqqq-backend.ddns.net:8001/uploads/orders/" + orderID + "/" + file.Filename[:len(file.Filename)-lenSymbols] + "pdf",
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
	}

	_, err = db.Conn.Model(&order).Set("pdf_link = ?", "http://fwqqq-backend.ddns.net:8001/uploads/orders/"+orderID+"/"+file.Filename[:len(file.Filename)-lenSymbols]+"pdf").Where("ID = ?", orderID).Update()
	if err != nil {
		return echo.NewHTTPError(http.StatusOK, err.Error())
	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully</p>", file.Filename))
}
