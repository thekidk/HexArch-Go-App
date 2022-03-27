package db

import (
	"database/sql"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driverName, dataSourceName string) (*Adapter, error) {

	//connect to db
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("DB connection error: %v", err)
	}

	// test db connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("DB ping failure: %v", err)
	}

	return &Adapter{db: db}, nil

}

func (da Adapter) CloseDBConnection() {

	err := da.db.Close()
	if err != nil {
		log.Fatalf("DB close failure: %v", err)
	}
	
}

func (da Adapter) AddToHistory(answer int32, operation string) error {

	queryString, args, err := sq.Insert("arith_history").Columns("date", "answer", "operation").Values(time.Now(),answer, operation).ToSql()
	if err != nil {
		 return err
	}

	_, err = da.db.Exec(queryString, args...)
	if err != nil {
		return err
	}

	return nil
}
