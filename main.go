/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 30-12-2020
 * |
 * | File Name:     main.go
 * +===============================================
 */

package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.File("/", "public/index.html")
	e.Static("/live", "docker/data")

	if err := e.Start("0.0.0.0:1378"); err != nil {
		fmt.Println(err)
	}
}
