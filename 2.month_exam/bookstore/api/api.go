package api

import (
	"encoding/json"
	"net/http"
	"store/internal/models"
	"store/internal/storage"
	"strconv"

	"github.com/gorilla/mux"
)

func AddNewBook(w http.ResponseWriter, r *http.Request) {
	DB, err := storage.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer DB.Close()
	var b models.Books
	err = json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = b.AddBook(DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(b)
	if err != nil {
		http.Error(w, "failed encoding", http.StatusInternalServerError)
		return
	}
}

func AddnewAuthor(w http.ResponseWriter, r *http.Request) {
	DB, err := storage.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer DB.Close()
	var a models.Author
	err = json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = a.AddAuthor(DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(a)
	if err != nil {
		http.Error(w, "failed encoding", http.StatusInternalServerError)
		return
	}
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	DB, err := storage.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer DB.Close()

	listBooks, err := models.GetBooks(DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	err = json.NewEncoder(w).Encode(listBooks)
	if err != nil {
		http.Error(w, "Error encoding listmovies", http.StatusInternalServerError)
		return
	}

}

func GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	DB, err := storage.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer DB.Close()
	listAuthors, err := models.GetAuthors(DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	err = json.NewEncoder(w).Encode(listAuthors)
	if err != nil {
		http.Error(w, "Error encoding listmovies", http.StatusInternalServerError)
		return
	}
}

func GetBooksById(w http.ResponseWriter, r *http.Request) {
	DB, err := storage.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer DB.Close()

	m := mux.Vars(r)
	idstr := m["id"]
	book_id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "Invalid id type:", http.StatusBadRequest)
		return
	}

	var book models.Books
	err = book.GetBookById(DB, book_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		http.Error(w, "error encoding movie data:", http.StatusInternalServerError)
		return
	}

}

func GetAuthorById(w http.ResponseWriter, r *http.Request) {
	DB, err := storage.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer DB.Close()

	m := mux.Vars(r)
	idstr := m["id"]
	author_id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "Invalid id type:", http.StatusBadRequest)
		return
	}

	var a models.Author
	err = a.GetAuthorById(DB, author_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	err = json.NewEncoder(w).Encode(a)
	if err != nil {
		http.Error(w, "error encoding movie data:", http.StatusInternalServerError)
		return
	}

}

func DeleteBookById(w http.ResponseWriter, r *http.Request) {
	DB, err := storage.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer DB.Close()

	m := mux.Vars(r)
	idStr := m["id"]
	book_id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "error: Invalid id format", http.StatusBadRequest)
		return
	}
	var book models.Books
	err = book.DeleteBook(DB, book_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	_, err = w.Write([]byte("book deleted succesfully !"))
	if err != nil {
		http.Error(w, "Error writing response:", http.StatusInternalServerError)
		return
	}

}

func DeleteAuthorById(w http.ResponseWriter, r *http.Request) {
	DB, err := storage.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer DB.Close()
	m := mux.Vars(r)
	idStr := m["id"]
	author_id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "error: Invalid id format", http.StatusBadRequest)
		return
	}
	var a models.Author
	err = a.DeleteAuthor(DB, author_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	_, err = w.Write([]byte("author deleted succesfully !"))
	if err != nil {
		http.Error(w, "Error writing response:", http.StatusInternalServerError)
		return
	}

}

func UpdateBookById(w http.ResponseWriter, r *http.Request) {
	DB, err := storage.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer DB.Close()
	m := mux.Vars(r)
	idStr := m["id"]
	book_id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "error: Invalid id.", http.StatusBadRequest)
		return
	}
	var b models.Books
	err = json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = b.UpdateBook(DB, book_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	err = json.NewEncoder(w).Encode(b)
	if err != nil {
		http.Error(w, "failed encoding json file", http.StatusInternalServerError)
		return
	}

}

func UpdateAuthorById(w http.ResponseWriter, r *http.Request) {
	DB, err := storage.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer DB.Close()
	m := mux.Vars(r)
	idStr := m["id"]
	author_id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "error: Invalid id.", http.StatusBadRequest)
		return
	}

	var a models.Author
	err = json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = a.UpdateAuthor(DB, author_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	err = json.NewEncoder(w).Encode(a)
	if err != nil {
		http.Error(w, "failed encoding json file", http.StatusInternalServerError)
		return
	}

}
