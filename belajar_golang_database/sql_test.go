package belajargolangdatabase

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// Query the database
	script := "INSERT INTO customer(id,name,email,balance,rating,birth_date,married) VALUES ('miftah', 'Miftah Az-zahra', 'miftahzahra@gmail.com', 500000, 76.4, '2004-06-15', false)"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	// Jika success insert customer
	fmt.Println("Success insert new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	// Menutup koneksi database
	defer db.Close()

	ctx := context.Background()

	// Query the database
	script := "SELECT * FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	// Menutup query scan tiap baris data from customer
	defer rows.Close()

	// Menampilkan query database
	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id   : ", id)
		fmt.Println("Name : ", name)
	}
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	// Menutup koneksi database
	defer db.Close()

	ctx := context.Background()

	// Query the database
	script := "SELECT id,name,email,balance,rating,birth_date,married,created_at FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	// Menutup query scan tiap baris data from customer
	defer rows.Close()

	// Menampilkan query database
	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate sql.NullTime
		var createdAt time.Time
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("================================================")
		fmt.Println("Id         : ", id)
		fmt.Println("Name       : ", name)
		if email.Valid {
			fmt.Println("Email      : ", email.String)
		}
		fmt.Println("Balance    : ", balance)
		fmt.Println("Rating     : ", rating)
		if birthDate.Valid {
			fmt.Println("Birth Date : ", birthDate.Time)
		}
		fmt.Println("Married    : ", married)
		fmt.Println("Created At : ", createdAt)
		fmt.Println("================================================")
	}
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	// Menutup koneksi database
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "salah"

	// Query the database
	script := "SELECT username FROM user where username = '" + username + "' AND password = '" + password + "' LIMIT 1"

	fmt.Println(script)
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	// Menutup query scan tiap baris data from customer
	defer rows.Close()

	// Menampilkan query database
	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Success login as", username)
	} else {
		fmt.Println("Failed Login to Application")
	}
}

func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	// Menutup koneksi database
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "admin"

	// Query the database
	script := "SELECT username FROM user where username = ? AND password = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}
	// Menutup query scan tiap baris data from customer
	defer rows.Close()

	// Menampilkan query database
	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Success login as", username)
	} else {
		fmt.Println("Failed Login to Application")
	}
}

func TestExecSqlSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "Rizal Pratama"
	password := "Rizalpr_099"

	// Query the database
	script := "INSERT INTO user(username, password) VALUES (?, ?)"
	_, err := db.ExecContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}

	// Jika success insert customer
	fmt.Println("Success insert new user")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "rizalpratama488@gmail.com"
	comment := "Aku benci orang munafik"

	// Query the database
	script := "INSERT INTO comments(email, comment) VALUES (?, ?)"
	result, err := db.ExecContext(ctx, script, email, comment)
	if err != nil {
		panic(err)
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	// Jika success insert customer
	fmt.Println("Success insert new comment with id ", insertId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	// Query the database
	script := "INSERT INTO comments(email, comment) VALUES (?,?)"
	// Prepare statement
	stmt, err := db.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	// Insert Data Sejumlah Input yang diinginkan
	for i := 0; i < 10; i++ {
		email := "rizal" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke-" + strconv.Itoa(i)
		result, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
		Id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Success insert new comment with id ", Id)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	// Query Database
	script := "INSERT INTO comments(email, comment) VALUES (?,?)"

	// Insert Data Sejumlah Input yang diinginkan
	for i := 0; i < 10; i++ {
		email := "rizal" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke-" + strconv.Itoa(i)
		result, err := tx.ExecContext(ctx, script, email, comment)
		if err != nil {
			panic(err)
		}
		Id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Success insert new comment with id ", Id)
	}
	tx.Rollback()
}
