package routes

import (
	"book-crud-project/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterBookRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBok).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
