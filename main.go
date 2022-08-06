package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type tableMahasiswa struct {
	db  *sql.DB
	ctx context.Context
}

func main() {
	db, err := sql.Open("mysql", "root:@/playground_database")
	ctx := context.Background()
	if err != nil {
		panic(err)
	}
	fmt.Println("connected to the database")
	defer db.Close()
	createtableMahasiswa := fmt.Sprintln(
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
	res, err := db.ExecContext(ctx, createtableMahasiswa)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
	dbMahasiswa := tableMahasiswa{db: db, ctx: ctx}
	res, err = dbMahasiswa.inputMahasiswa("aditya ramadhan", "215150201111543007", "adityarrr@gmail.com", "Teknik Informatika", "FILKOM")
	if err != nil {
		panic(err)
	}
	fmt.Println(res.LastInsertId())
}

func (t *tableMahasiswa) inputMahasiswa(nama, nim, email, jurusan, fakultas string) (sql.Result, error) {
	inputMahasiswa := fmt.Sprintf(`
	INSERT INTO mahasiswa(nama, nim, email, jurusan, fakultas)
	VALUES("%s", "%s", "%s", "%s", "%s")`, nama, nim, email, jurusan, fakultas)
	res, err := t.db.ExecContext(t.ctx, inputMahasiswa)
	if err != nil {
		return nil, err
	}
	return res, nil

}
