package main

import (
	"fmt"
	"net/http"
	"os"
	"store/api"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func main() {
	r := mux.NewRouter()
	// err :=godotenv.Load(".env")
	// if err !=nil {
	// 	log.Fatal(err)
	// }

	r.HandleFunc("/authors", api.GetAllAuthors).Methods("GET")
	r.Handle("/author/{id}/", api.IsValidId(http.HandlerFunc(api.GetAuthorById))).Methods("GET")
	r.Handle("/author", api.ContentTypeMiddle(http.HandlerFunc(api.AddnewAuthor))).Methods("POST")
	r.Handle("/author/{id}/", api.CheckPermission(http.HandlerFunc(api.DeleteAuthorById))).Methods("DELETE")
	r.Handle("/author/{id}/", api.CheckPermission(http.HandlerFunc(api.UpdateAuthorById))).Methods("PUT")

	r.HandleFunc("/books", api.GetAllBooks).Methods("GET")
	r.Handle("/book/{id}/", api.IsValidId(http.HandlerFunc(api.GetBooksById))).Methods("GET")
	r.Handle("/book", api.ContentTypeMiddle(http.HandlerFunc(api.AddNewBook))).Methods("POST")
	r.Handle("/book/{id}/", api.CheckPermission(http.HandlerFunc(api.DeleteBookById))).Methods("DELETE")
	r.Handle("/book/{id}/", api.CheckPermission(http.HandlerFunc(api.UpdateBookById))).Methods("PUT")

	fmt.Printf("Server is listening on port %s\n", os.Getenv("SERVER_HOST") + os.Getenv("SERVER_PORT"))
	http.ListenAndServe(os.Getenv("SERVER_PORT"), r)
}
