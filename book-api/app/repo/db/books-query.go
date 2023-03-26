package db

import (
	"book-api/app/entity"
	"book-api/pkg/database"
	"database/sql"
)

func GetAllData() ([]entity.Book, error) {

	var hasil = []entity.Book{}

	db, err := database.Connect()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	sql := "SELECT * FROM books"

	rows, err := db.Query(sql)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var book = entity.Book{}

		err = rows.Scan(&book.BookID, &book.Title, &book.Author, &book.Description)
		if err != nil {
			return nil, err
		}
		hasil = append(hasil, book)
	}
	return hasil, nil

}

func GetDataByID(idparam int) ([]entity.Book, error) {

	var hasil = []entity.Book{}

	db, err := database.Connect()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	sqlQuery := "SELECT * FROM books WHERE id_book=$1"
	var book = entity.Book{}
	err = db.QueryRow(sqlQuery, idparam).Scan(&book.BookID, &book.Title, &book.Author, &book.Description)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	hasil = append(hasil, book)

	return hasil, nil

}

func InsertData(book entity.Book) (hasil string, err error) {
	db, err := database.Connect()
	if err != nil {
		return "error", err
	}

	defer db.Close()

	sql := "INSERT INTO books(title, author, description) VALUES ($1, $2, $3)"
	_, err = db.Exec(sql, book.Title, book.Author, book.Description)

	if err != nil {
		return "error", err
	}

	return "Data berhasil di-insert", nil
}

func UpdateData(idparam int, book entity.Book) (hasil string, err error) {
	db, err := database.Connect()
	if err != nil {
		return "error", err
	}

	defer db.Close()

	sql := "UPDATE books SET title=$1, author=$2, description=$3 WHERE id_book=$4"
	_, err = db.Exec(sql, book.Title, book.Author, book.Description, idparam)

	if err != nil {
		return "error", err
	}

	return "Data buku berhasil di-update", nil
}

func DeleteData(idparam int) (string, error) {
	db, err := database.Connect()
	if err != nil {
		return "error", err
	}

	defer db.Close()

	sql := "DELETE FROM books WHERE id_book=$1"
	_, err = db.Exec(sql, idparam)

	if err != nil {
		return "error", err
	}

	return "Data buku berhasil dihapus", nil

}
