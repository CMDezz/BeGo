package queries

import "github.com/jmoiron/sqlx"

type Queries struct{}

type IQueries interface {
	QueriesTestData(db *sqlx.DB)
}

func NewQueries() IQueries {
	return &Queries{}
}
