package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@/playground_database")
	ctx := context.Background()
	if err != nil {
		panic(err)
	}
	fmt.Println("connected to the database")
	defer db.Close()
	tableMahasiswa := fmt.Sprintln(
		`CREATE TABLE IF NOT EXISTS mahasiswa (
			id INTEGER PRIMARY KEY AUTO_INCREMENT,
			nama TEXT NOT NULL,
			nim TEXT NOT NULL UNIQUE,
			email TEXT NOT NULL UNIQUE,
			jurusan TEXT NOT NULL,
			fakultas TEXT NOT NULL,
			is_active BOOLEAN DEFAULT 1,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
			)`)
	res, err := db.ExecContext(ctx, tableMahasiswa)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
	res, err = inputMahasiswa("aditya rizky", "215150201111007", "aditya@gmail.com", "Teknik Informatika", "FILKOM", db, ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func inputMahasiswa(nama, nim, email, jurusan, fakultas string, db *sql.DB, ctx context.Context) (sql.Result, error) {
	inputMahasiswa := fmt.Sprintf(`
	INSERT INTO mahasiswa(nama, nim, email, jurusan, fakultas)
	VALUES("%s", "%s", "%s", "%s", "%s")`, nama, nim, email, jurusan, fakultas)
	res, err := db.ExecContext(ctx, inputMahasiswa)
	if err != nil {
		return nil, err
	}
	return res, nil

}
