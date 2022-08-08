package table_matkul_nilai

import (
	"context"
	"database/sql"
	"fmt"
)

type tableMatkulNilai struct {
	db  *sql.DB
	ctx context.Context
}

func NewTableMatkulNilaiMahasiswa(db *sql.DB, ctx context.Context) (*tableMatkulNilai, error) {
	makeTable := fmt.Sprintln(
		`CREATE TABLE IF NOT EXISTS matkul_nilai (
			id INT AUTO_INCREMENT,
			nilai INT NOT NULL,
			mahasiswa_id INT NOT NULL,
			mata_kuliah TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			PRIMARY KEY (id),
			FOREIGN KEY (mahasiswa_id) REFERENCES mahasiswa(id)
			)`)
	_, err := db.ExecContext(ctx, makeTable)
	if err != nil {
		return nil, err
	}
	return &tableMatkulNilai{db: db, ctx: ctx}, nil
}

func (t *tableMatkulNilai) InputMatkulNilai(nilai int, mahasiswaID int, mataKuliah string) (sql.Result, error) {
	inputMatkulNilai := fmt.Sprintf(`
	INSERT INTO matkul_nilai(nilai, mahasiswa_id, mata_kuliah)
	VALUES(%d, %d, "%s")`, nilai, mahasiswaID, mataKuliah)
	res, err := t.db.ExecContext(t.ctx, inputMatkulNilai)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (t *tableMatkulNilai) UpdateNilai(id int, nilai int) (sql.Result, error) {
	updateNilai := fmt.Sprintf(`
	UPDATE matkul_nilai
	SET nilai = %d
	WHERE id = %d`, nilai, id)
	res, err := t.db.ExecContext(t.ctx, updateNilai)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type ScanTableMatkulNilai struct {
	Id          int    `json:"id"`
	Nilai       int    `db:"nilai"`
	MahasiswaID int    `db:"mahasiswa_id"`
	MataKuliah  string `db:"mata_kuliah"`
	CreatedAt   string `db:"created_at"`
	UpdatedAt   string `db:"updated_at"`
}

func (t *tableMatkulNilai) ReadByIdMahasiswa(id int) (*ScanTableMatkulNilai, error) {
	readByIdMahasiswa := fmt.Sprintf(`
	SELECT * FROM matkul_nilai
	WHERE mahasiswa_id = %d`, id)
	rows, err := t.db.QueryContext(t.ctx, readByIdMahasiswa)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var scanTableMatkulNilai ScanTableMatkulNilai
	for rows.Next() {
		err := rows.Scan(&scanTableMatkulNilai.Id, &scanTableMatkulNilai.Nilai, &scanTableMatkulNilai.MahasiswaID, &scanTableMatkulNilai.MataKuliah, &scanTableMatkulNilai.CreatedAt, &scanTableMatkulNilai.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}
	return &scanTableMatkulNilai, nil
}
