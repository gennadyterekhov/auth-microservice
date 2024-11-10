package models

import (
	"database/sql"
)

type User struct {
	ID       int64
	Login    string
	Password string
}

func (user *User) ScanRow(row *sql.Row) error {
	return row.Scan(&(user.ID), &(user.Login), &(user.Password))
}
