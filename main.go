package main

import (
	tablemahasiswa "app_basis_data/table_mahasiswa"
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
	mahasiswa, err := tablemahasiswa.NewTableMahasiswa(db, ctx)
	if err != nil {
		panic(err)
	}
	_, err = mahasiswa.InputMahasiswa("Rizki", "12345", "rizky@gmail.com", "Teknik Informatika", "FILKOM")
	if err != nil {
		panic(err)
	}
	_, err = mahasiswa.UpdateMahasiswaIsActive("12345", false)
	if err != nil {
		panic(err)
	}
	fmt.Println("data inserted with status mahasiswa.is_active = false")
}
