package models

import (
	"database/sql"
	"log"
)

type DBModel struct {
	DB *sql.DB
	infoLog *log.Logger
	errorLog *log.Logger
}

func NewDBModel(db *sql.DB, infoLog, errorLog *log.Logger) *DBModel {
	return &DBModel{
		DB: db,
		infoLog: infoLog,
		errorLog: errorLog,
	}
}