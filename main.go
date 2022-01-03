package main

import (
	"fmt"
	"github.com/iris/web/controllers"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	//注册模版目录
	app.RegisterView(iris.HTML("./web/views", ".html"))
	mvc.New(app.Party("/hello")).Handle(new(controllers.MovieController))
	err := app.Run(iris.Addr("localhost:8080"))
	if err != nil {
		fmt.Println(err)
	}
}
