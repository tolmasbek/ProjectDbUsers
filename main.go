package main

import (
	"fmt"
	"database/sql"
	"go/database"
	_"sqlite/go-sqlite3"
	"log"
	"go/pkg/accounts"
	"go/pkg/users"
	// "go/pkg/currency"
)

func main() {

	db, err := sql.Open("sqlite3", "ProjectUsersDb.db")
	if err != nil {
		log.Fatal("Data Base conection is failed", err)
	}
	err = database.DbCreateUsers(db)
	if err != nil {
		fmt.Println("not found", err)
	}



	open, err1 := sql.Open("sqlite3", "ProjectUsersDb.db")
	if err1 != nil {
		log.Fatal("not found", err1)
	}
	err1 = database.DbCreateCurr(open)
	if err1 != nil {
		fmt.Println("not found", err1, )
	}



	db, err2 := sql.Open("sqlite3", "Bank")
	if err2 != nil {
		log.Fatal("not found", err2)
	}
	err2 = database.DbCreateAcc(db)
	if err2 != nil {
		fmt.Println("not found", err2)

	}


	acc := accounts.Accounts{
		Id:       	1,
		UserId:   	1,
		Numb:   	2017050131,
		Amount:   	1000,
		Currency: 	"TJS",
	}
	user := users.Users{
		Id:        	1,
		FirstName: 	"Suhrob",
		LastName: 	"Farmonov",
		Age: 		20,
		Gender:    	"male",
		Login:     	"ss9999",
		Password:  	"999999",
	}



	err3 := AddNewClient(db, user)
	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Println(user)

	err4 := AddNewAcount(db, acc)
	if err4 != nil {
		fmt.Println(err4)
	}
	fmt.Println(acc)

	err5 := RemoveById(db)
	if err5 != nil {
		fmt.Println(err5)
	}
	fmt.Println("Sorry I can not remove")

	err6 := Update(db)
	if err6 != nil {
		fmt.Println(err6)
	}
	fmt.Println("Sorry nothing update")
}

func AddNewClient(db *sql.DB, client users.Users) (err3 error) {

	_, err3 = db.Exec("Insert Into Users(FirstName, LastName, Age, Gender, Login, Password) values (($1), ($2), ($3), ($4), ($5), ($6))", client.FirstName, client.LastName, client.Age, client.Gender, client.Login, client.Password)
	if err3 != nil {
		fmt.Println(err3)
	}
	return err3
}

func AddNewAcount(db *sql.DB, acount accounts.Accounts) (err4 error)  {
	_, err4 = db.Exec("Insert Into Accounts(UsersId, Numb, Amount, Currency) values (($1), ($2), ($3), ($4))", acount.UserId, acount.Numb, acount.Amount, acount.Currency)
	if err4 != nil {
		fmt.Println(err4)
	}
	return err4
}

func RemoveById(db *sql.DB) (err5 error)  {
	_, err5 = db.Exec("Delete From Users Where Id = 1")
	if err5 != nil {
		fmt.Println(err5)
	}
	return err5
}

func Update(db *sql.DB) (err6 error) {
	_, err6 = db.Exec("Update Account Set Amount = Amount + 33000 ")
	if err6 != nil {
		fmt.Println(err6)
	}
	return err6
}

