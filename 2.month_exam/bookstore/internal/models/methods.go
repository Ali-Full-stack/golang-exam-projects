package models

import (
	"database/sql"
)

func (b *Books) AddBook(db *sql.DB) error {
	query := `
		insert into books(
			title,category,author_id,publication_date,isbn,description,created_at,updated_at
		)
		values(
			$1, $2, $3, $4, $5, $6, $7, $8
		)
		returning
		book_id
	`

	row := db.QueryRow(query, b.Title, b.Category, b.Author_id, b.Publication_date, b.Isbn, b.Description, b.Created_at, b.Updated_at)
	if err := row.Scan(&b.Book_id); err != nil {
		return err
	}
	return nil
}

func (a *Author) AddAuthor(db *sql.DB) error {
	query := `
	insert into author(
		name,birth_date,biography,created_at,updated_at
	)
	values(
		$1, $2, $3, $4, $5
	)returning
	author_id
`
	row := db.QueryRow(query, a.Name, a.Birth_date, a.Biography, a.Created_at, a.Updated_at)
	if err := row.Scan(&a.Author_id); err != nil {
		return err
	}

	return nil
}

func GetBooks(db *sql.DB) ([]Books, error) {
	query := `
		select book_id, title, category, author_id, publication_date,isbn, description, created_at,updated_at
		from books
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var listBooks []Books

	for rows.Next() {
		var b Books
		err = rows.Scan(&b.Book_id, &b.Title, &b.Category, &b.Author_id, &b.Publication_date, &b.Isbn, &b.Description, &b.Created_at, &b.Updated_at)
		if err != nil {
			return nil, err
		}
		listBooks = append(listBooks, b)
	}
	return listBooks, nil

}

func GetAuthors(db *sql.DB) ([]Author, error) {
	query := `
		select  author_id, name, birth_date, biography,created_at,updated_at
		from author
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var listAuthors []Author

	for rows.Next() {
		var a Author
		err = rows.Scan(&a.Author_id, &a.Name, &a.Birth_date, &a.Biography, &a.Created_at, &a.Updated_at)
		if err != nil {
			return nil, err
		}

		listAuthors = append(listAuthors, a)
	}
	return listAuthors, nil
}

func (b *Books) GetBookById(db *sql.DB, id int) error {

	query := `
	select title,category,author_id,publication_date,isbn,description,created_at,updated_at
	from books
	where book_id = $1
	`
	row := db.QueryRow(query, id)
	err := row.Scan(&b.Title, &b.Category, &b.Author_id, &b.Publication_date, &b.Isbn, &b.Description, &b.Created_at, &b.Updated_at)
	if err != nil {
		if err == sql.ErrNoRows {
			return err
		}
		return err
	}
	b.Book_id = id
	return nil
}

func (a *Author) GetAuthorById(db *sql.DB, id int) error {
	query := `
	select author_id,name,birth_date,biography,created_at,updated_at
	from author
	where author_id = $1
	`
	row := db.QueryRow(query, id)
	err := row.Scan(&a.Author_id, &a.Name, &a.Birth_date, &a.Biography, &a.Created_at, &a.Updated_at)
	if err != nil {
		if err == sql.ErrNoRows {
			return err
		}
		return err
	}
	a.Author_id = id
	return nil

}

func (b *Books) DeleteBook(db *sql.DB, id int)error {
	query := `
		delete from books
		where book_id = $1
	`

	res, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return err
	}
	return nil
}

func (a *Author) DeleteAuthor(db *sql.DB, id int)error{

	query := `
		delete from author
		where author_id = $1
	`

	res, err := db.Exec(query,id)
	if err != nil {
		return err
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return err
	}

	return nil
}

func (b *Books) UpdateBook(db *sql.DB, id int)error{
	query := `
		update books
		set title=$1, category=$2, author_id=$3, publication_date=$4, isbn=$5, description=$6, created_at=$7
		where book_id=$8
	`

	res, err := db.Exec(query, b.Title, b.Category, b.Author_id, b.Publication_date, b.Isbn, b.Description, b.Created_at, id)
	if err != nil {
		return err
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return nil
	}
	b.Updated_at =CurrentStringtime()
	return	nil
}

func (a *Author) UpdateAuthor(db *sql.DB, id int)error{
	query := `
		update author
		set name=$1, birth_date=$2, biography=$3, created_at=$4
		where author_id=$5
	`

	res, err := db.Exec(query, a.Name, a.Birth_date, a.Biography, a.Created_at, id)
	if err != nil {
		return err
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return err
	}
	a.Updated_at =CurrentStringtime()
	return nil
}
