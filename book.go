package main

import(
	"fmt"
	"strings"
)

type BookReference struct{
	firstName string
	lastName string
	bookTitle string
	placeOfPublication string
	publishingCompany string
	year string
}

//GetBookInfo is used to interact with the user and get information about the book
func GetBookInfo() string{
	var book BookReference
	fmt.Println("Author's first name:")
	fmt.Scanf("%s", &book.firstName)
	fmt.Println("Author's last name:")
	fmt.Scanf("%s", &book.lastName)
	fmt.Println("Book title:")
	fmt.Scanf("%s", &book.bookTitle)
	fmt.Println("Place of publication:")
	fmt.Scanf("%s", &book.placeOfPublication)
	fmt.Println("Publishing company:")
	fmt.Scanf("%s", &book.publishingCompany)
	fmt.Println("Year of publication:")
	fmt.Scanf("%s", &book.year)
	return ReturnBookReference(&book)
}

//ReturnBookReference is used to concatenate all information given in one string and print it
func ReturnBookReference(b * BookReference) string{
	var template string
	b.lastName= strings.ToUpper(b.lastName)
	template = fmt.Sprintf("%s, %s. %s. %s: %s, %s.", b.lastName, b.firstName, b.bookTitle, b.placeOfPublication, b.publishingCompany,  b.year)
	return template
}






