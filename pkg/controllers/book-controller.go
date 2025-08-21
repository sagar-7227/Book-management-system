package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/sagar-7227/Book-management-system/pkg/models"
	"github.com/sagar-7227/Book-management-system/pkg/utils"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request){
	newbooks:= models.GetAllBooks()
	res, _ := json.Marshal(newbooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing book ID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	bookDetails, _ := models.GetBookById(id)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {

	CreateBook := &models.Book{}
	utils.ParseRequestBody(r, &CreateBook)
	newBook := CreateBook.CreateBook()
	res, _ := json.Marshal(newBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing book ID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	bookdelete := models.DeleteBook(id)
	res, _ := json.Marshal(bookdelete)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing book ID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	BookDetails, _ := models.GetBookById(id)
	if BookDetails == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	utils.ParseRequestBody(r, BookDetails)
	updatedBook := BookDetails.UpdateBook()
	res, _ := json.Marshal(updatedBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}