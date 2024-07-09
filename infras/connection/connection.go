package connection

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

func InitConnection(DbDriver string, DbSource string) (*sql.DB, *sqlx.DB) {
	//SQL
	sqlContext, err := sql.Open(DbDriver, DbSource)
	fmt.Println(sqlContext.Ping())
	if err != nil || sqlContext.Ping() != nil {
		fmt.Println(err)

		log.Fatal("Internal server error: Cannot init SQL connection")
	}

	//SQLX
	sqlxContext, err := sqlx.Open(DbDriver, DbSource)
	fmt.Println(sqlxContext.Ping())

	if err != nil || sqlxContext.Ping() != nil {
		fmt.Println(err)

		log.Fatal("Internal server error: Cannot init SQLX connection")
	}

	return sqlContext, sqlxContext
}
