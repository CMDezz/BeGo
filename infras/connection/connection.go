package connection

import (
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
)

func InitConnection(DbDriver string, DbSource string) (*sql.DB, *sqlx.DB) {
	//SQL
	sqlContext, err := sql.Open(DbDriver, DbSource)
	if err != nil || sqlContext.Ping() != nil {
		log.Fatal("Internal server error: Cannot init SQL connection")
	}

	//SQLX
	sqlxContext, err := sqlx.Open(DbDriver, DbSource)
	if err != nil || sqlxContext.Ping() != nil {
		log.Fatal("Internal server error: Cannot init SQL connection")
	}

	return sqlContext, sqlxContext
}
