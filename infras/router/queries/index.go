package queries

import "github.com/jmoiron/sqlx"

type Queries struct {
	db *sqlx.DB
}

type IQueries interface {
	IPostQueries
	IUserQueries
}

func NewQueries(db *sqlx.DB) IQueries {
	return &Queries{db: db}
}
