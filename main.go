package main

import (
	tablemahasiswa "app_basis_data/table_mahasiswa"
	"app_basis_data/table_matkul_nilai"
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
	_, err = table_matkul_nilai.NewTableMatkulNilaiMahasiswa(db, ctx)
	if err != nil {
		panic(err)
	}
	// nilai.InputMatkulNilai(90, 1, "Algoritma dan Struktur Data")
	// _, err = mahasiswa.InputMahasiswa("Rizki", "12345", "rizky@gmail.com", "Teknik Informatika", "FILKOM")
	// if err != nil {
	// 	panic(err)
	// }
	// _, err = mahasiswa.UpdateMahasiswaIsActive("12345", false)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("data inserted with status mahasiswa.is_active = false")\
	dataMahasiswa, err := mahasiswa.ReadMahasiswaByIdAndMatkulNilai(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(dataMahasiswa)
	/*
		connected to the database
		[{{1 Aditya Rizky 123456789 aditya@gmail.com TIF FILKOM true 2022-08-07 16:30:15 2022-08-07 16:30:15} {1 90 1 Pemrograman Dasar 2022-08-08 22:39:09 2022-08-08 22:39:09}} {{1 Aditya Rizky 123456789 aditya@gmail.com TIF FILKOM true 2022-08-07 16:30:15 2022-08-07 16:30:15} {2 90 1 Algoritma dan Struktur Data 2022-08-08 22:40:25 2022-08-08 22:40:25}}]

	*/
}
