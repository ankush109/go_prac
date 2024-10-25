package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Book struct {
	Title  string
	Author string
}

type Todo struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
}

type LibraryController struct {
}

func (c LibraryController) GetBook(title string) (*Book, error) {
	if title == "a" {
		return &Book{Title: "Tails of Gur Das", Author: "Ankush Banerjee"}, nil
	}
	return nil, fmt.Errorf("Book not found ")
}

func main() {
	// l := LibraryController{}
	// book, err := l.GetBook("a")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("Retrieved book is %s by %s\n", book.Title, book.Author)
	url := "https://jsonplaceholder.typicode.com/todos/1"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("error in reponse", err)
		return
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	var todo Todo
	err = json.Unmarshal(body, &todo)
	if err != nil {
		fmt.Println("err decoding the reponse", err)
		return
	}
	println(todo.Title)
}
