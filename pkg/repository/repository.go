package repository

type Authorization interface {
}

type BrawlersList interface {
}
type BrawlersItem interface {
}
type Repository interface {
	Authorization
	BrawlersList
	BrawlersItem
}

func NewRepository() *Repository {
	return &Repository{}
}
