package database

import (
	"fmt"
	"database/sql"
	"ProjectUsers/pkg/accounts"
)

func DbCreateUsers(db*sql.DB)(err error){
	const UserDb = `CREATE TABLE Users (
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
		FirstName TEXT NOT NULL,
		LastName TEXT NOT NULL,
		Age INTEGER NOT NULL,
		Gender TEXT NOT NULL,
		Login TEXT NOT NULL,
		Password  TEXT NOT NULL,
)`
	_, err = db.Exec(UserDb)
	if err != nil {
		return nil
	}
	return nil
}
func DbCreateCurr(db *sql.DB)(err error){
	const CurrencuDb = `CREATE TABLE Currency(
		Id INTEGER PRIMARY KEY AUTOINCREMENT
		Name TEXT,
)`
	_, err = db.Exec(CurrencuDb)
	if err != nil {
		return nil
	}
	return nil
}
func DbCreateAcc(db *sql.DB)(err error){
	const AccountsDb = `CREATE TABLE IF NOT EXIST AccountsDb(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	UserId INTEGER REFERENCES Users(Id) NOT NULL,
	Number INTEGER,
	Amount INTEGER,
	Currency INTEGER REFERENCES Currency(Id),
	REMOVE BOOLEAN NOT NULL DEFAULT FALSE
	)`
	_, err = db.Exec(AccountsDb)
	if err != nil {
		return nil
	}
	return nil
}