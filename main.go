package main

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo"

	"httpstatus/statuscodes"
)

func main() {
	e := echo.New()

	e.GET("/:statusCode", func(c echo.Context) error {
		allowedCodes := statuscodes.GetAllowedCodes()
		statusCode, _ := strconv.Atoi(c.Param("statusCode"))

		if allowedCodes.IsInList(statusCode) {
			return c.String(statusCode, "")
		}

		return c.String(500, fmt.Sprintf("Error: status code %s is not available.", c.Param("statusCode")))
	})

	e.Logger.Fatal(e.Start(":8080"))
}
