package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World")
}

func ParamsEcho(c echo.Context) error {
	val := c.Param("search")
	return c.String(http.StatusOK, val)
}

func QueryEcho(c echo.Context) error {
	val := c.QueryParam("name")
	fmt.Println(val)
	return c.String(http.StatusOK, val)
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`

}

func UserRegister(c echo.Context) error{
	var u User
	err := c.Bind(&u);
	if(err!=nil){
		return err
	}

	fmt.Println(u)

	return c.JSON(http.StatusOK,u)
}

func main() {
	e := echo.New()

	e.GET("/", HelloWorld)
	e.GET("/:search", ParamsEcho)
	e.GET("/query", QueryEcho)

	e.POST("/user",UserRegister)

	e.Logger.Fatal(e.Start(":1323"))
}
