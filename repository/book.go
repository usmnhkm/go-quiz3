package repository

import (
	"database/sql"
	"quiz3/structs"
)

func GetAllBooks(db *sql.DB) (books []structs.Book, err error) {
	sql := "select * from book"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var book structs.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Description, &book.ImageUrl, &book.ReleaseYear, &book.Price,
			&book.TotalPage, &book.Thickness, &book.CreatedAt, &book.UpdatedAt, &book.CategoryID)
		if err != nil {
			panic(err)
		}
		books = append(books, book)
	}
	return
}

func InsertBook(db *sql.DB, book structs.Book) (err error) {
	sql := "insert into book (title, description, image_url, release_year, price, total_page, thickness, category_id) values ($1, $2, $3, $4, $5, $6, $7, $8)"
	errs := db.QueryRow(sql, book.Title, book.Description, book.ImageUrl, book.ReleaseYear, book.Price, book.TotalPage, book.Thickness, book.CategoryID)
	return errs.Err()
}

func UpdateBook(db *sql.DB, book structs.Book) (err error) {
	sql := "update book set title=$1, description=$2, image_url=$3, release_year=$4, price=$5, total_page=$6, thickness=$7, category_id=$8 where id = $9"
	errs := db.QueryRow(sql, book.Title, book.Description, book.ImageUrl, book.ReleaseYear, book.Price, book.TotalPage, book.Thickness, book.CategoryID, book.ID)
	return errs.Err()
}

func DeleteBook(db *sql.DB, book structs.Book) (err error) {
	sql := "delete from book where id=$1"
	errs := db.QueryRow(sql, book.ID)
	return errs.Err()
}
