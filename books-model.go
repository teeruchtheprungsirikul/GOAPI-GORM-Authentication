package main

import (
	"log"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
}

func createBook(db *gorm.DB, book *Book) error {
	result := db.Create(book)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func getBook(db *gorm.DB, id int) *Book {
	var book Book
	result := db.First(&book, id)

	if result.Error != nil {
		log.Fatalf("Error get book: %v", result.Error)
	}

	return &book
}

func getBooks(db *gorm.DB) []Book {
	var books []Book
	result := db.Find(&books)

	if result.Error != nil {
		log.Fatalf("Error get book: %v", result.Error)
	}

	return books
}

func updateBook(db *gorm.DB, book *Book) error {
	result := db.Model(&book).Updates(book)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func deleteBook(db *gorm.DB, id int) error {
	var book Book
	// Hard Delete
	result := db.Unscoped().Delete(&book, id)
	/*
		Soft Delete
		result := db.Delete(&book, id)
	*/

	if result.Error != nil {
		return result.Error
	}

	return nil

}

func searchBook(db *gorm.DB, bookname string) []Book {
	var books []Book

	result := db.Where("name = ?", bookname).Order("price desc").Find(&books)
	if result.Error != nil {
		log.Fatalf("Search book failed: %v", result.Error)
	}

	return books
}
