package tablemahasiswa

import (
	"context"
	"database/sql"
	"fmt"
)

type tableMahasiswa struct {
	db  *sql.DB
	ctx context.Context
}

func NewTableMahasiswa(db *sql.DB, ctx context.Context) (*tableMahasiswa, error) {
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
	_, err := db.ExecContext(ctx, createtableMahasiswa)
	if err != nil {
		return nil, err
	}
	return &tableMahasiswa{db: db, ctx: ctx}, nil
}

func (t *tableMahasiswa) InputMahasiswa(nama, nim, email, jurusan, fakultas string) (sql.Result, error) {
	inputMahasiswa := fmt.Sprintf(`
	INSERT INTO mahasiswa(nama, nim, email, jurusan, fakultas)
	VALUES("%s", "%s", "%s", "%s", "%s")`, nama, nim, email, jurusan, fakultas)
	res, err := t.db.ExecContext(t.ctx, inputMahasiswa)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (t *tableMahasiswa) UpdateMahasiswaIsActive(nim string, isActive bool) (sql.Result, error) {
	updateMahasiswa := fmt.Sprintf(`
	UPDATE mahasiswa
	SET is_active = %t
	WHERE nim = "%s"`, isActive, nim)
	res, err := t.db.ExecContext(t.ctx, updateMahasiswa)
	if err != nil {
		return nil, err
	}
	return res, nil
}
