package controllers

import (
	"github.com/iris/repositories"
	"github.com/iris/services"
	"github.com/kataras/iris/v12/mvc"
)

type MovieController struct {
}

func (c *MovieController) Get() mvc.View {
	movieRepository := repositories.NewMovieManager()
	movieService := services.NewMovieServiceManager(movieRepository)
	MovieResult := movieService.ShowMovieName()
	return mvc.View{
		Name: "movie/index.html",
		Data: MovieResult,
	}
}
