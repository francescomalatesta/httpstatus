package main

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo"

	"httpstatus/statuscodes"
)

func handleStatusCodeRequest(c echo.Context) error {
	allowedCodes := statuscodes.GetAllowedCodes()
	statusCode, _ := strconv.Atoi(c.Param("statusCode"))

	if allowedCodes.IsInList(statusCode) {
		return c.String(statusCode, "")
	}

	return c.String(500, fmt.Sprintf("Error: status code %s is not available.", c.Param("statusCode")))
}

func main() {
	e := echo.New()

	e.File("/", "public/index.html")

	e.GET("/:statusCode", handleStatusCodeRequest)
	e.GET("/:statusCode/*", handleStatusCodeRequest)
	e.POST("/:statusCode", handleStatusCodeRequest)
	e.POST("/:statusCode/*", handleStatusCodeRequest)
	e.PUT("/:statusCode", handleStatusCodeRequest)
	e.PUT("/:statusCode/*", handleStatusCodeRequest)
	e.PATCH("/:statusCode", handleStatusCodeRequest)
	e.PATCH("/:statusCode/*", handleStatusCodeRequest)
	e.DELETE("/:statusCode", handleStatusCodeRequest)
	e.DELETE("/:statusCode/*", handleStatusCodeRequest)
	e.OPTIONS("/:statusCode", handleStatusCodeRequest)
	e.OPTIONS("/:statusCode/*", handleStatusCodeRequest)
	e.HEAD("/:statusCode", handleStatusCodeRequest)
	e.HEAD("/:statusCode/*", handleStatusCodeRequest)

	e.Logger.Fatal(e.Start(":8080"))
}
