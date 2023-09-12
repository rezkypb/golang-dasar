package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "rezky",
		"address": "gorontalo",
	}
	person["title"] = "Programmer"
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book map[string]string = make(map[string]string)

	book["title"] = "Belajar golang"
	book["author"] = "Rezky"
	book["ups"] = "salah"
	fmt.Println(book)
	delete(book, "ups")
	fmt.Println(book)
}

