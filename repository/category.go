package repository

import (
	"database/sql"
	"quiz3/structs"
)

func GetAllCategories(db *sql.DB) (categories []structs.Category, err error) {
	sql := "select * from category"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var category structs.Category
		err := rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			panic(err)
		}
		categories = append(categories, category)
	}
	return
}

func GetBooksWithCategory(db *sql.DB, category_id int) (books []structs.Book, err error) {
	sql := "select * from book where category_id=$1"

	rows, err := db.Query(sql, category_id)
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

func InsertCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "insert into category (name) values ($1)"
	errs := db.QueryRow(sql, category.Name)
	return errs.Err()
}

func UpdateCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "update category set name=$1 where id=$2"
	errs := db.QueryRow(sql, category.Name, category.ID)
	return errs.Err()
}

func DeleteCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "delete from category where id=$1"
	errs := db.QueryRow(sql, category.ID)
	return errs.Err()
}
