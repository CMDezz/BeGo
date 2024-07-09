package queries

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func (queries *Queries) QueriesTestData(db *sqlx.DB) {
	fmt.Println("QueriesTestData")
}
