package database

import (
	pkg "book-api/pkg/common"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", pkg.Host, pkg.Port, pkg.User, pkg.Password, pkg.DBname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		return nil, err
	}

	return db, nil

}
