package repositories

import "github.com/iris/datamodels"

type MovieRepository interface {
	GetMovieName() string
}

type MovieManager struct {
}

func NewMovieManager() MovieRepository {
	return &MovieManager{}
}

func (m *MovieManager) GetMovieName() string {
	//todo:去数据库查询 等到Name
	//给模型赋值
	movie := &datamodels.Movie{Name: "黑客地瓜"}
	return movie.Name
}
