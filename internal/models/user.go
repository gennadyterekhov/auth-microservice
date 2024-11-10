package models

import (
	"database/sql"
)

type User struct {
	ID       int64
	Login    string
	Password string
}

// ScanRow Implement ScanRow for User. see abstract.Scannable
func (user *User) ScanRow(row *sql.Row) error {
	if user == nil {
		user = new(User)
	}
	return row.Scan(&(user.ID), &(user.Login), &(user.Password))
}

// ScanRows Implement ScanRows for User. see abstract.Scannable
func (user *User) ScanRows(rows *sql.Rows) error {
	if user == nil {
		user = new(User)
	}
	return rows.Scan(&(user.ID), &(user.Login), &(user.Password))
}
