package main

import(
	"fmt"
	"strings"
)

type websiteReference struct{
	firstName string
	lastName string
	articleTitle string
	websiteName string
	url string
	year string
	dayOfAccess string
	monthOfAccess string
	yearOfAcess string
}

func GetWebsiteInfo() string{
	var website websiteReference
	fmt.Println("Author's first name:")
	fmt.Scanf("%s", &website.firstName)
	fmt.Println("Author's last name:")
	fmt.Scanf("%s", &website.lastName)
	fmt.Println("Article title:")
	fmt.Scanf("%s", &website.articleTitle)
	fmt.Println("Website name publication:")
	fmt.Scanf("%s", &website.websiteName)
	fmt.Println("URL:")
	fmt.Scanf("%s", &website.url)
	fmt.Println("Year of publication:")
	fmt.Scanf("%s", &website.year)
	fmt.Println("Day of access:")
	fmt.Scanf("%s", &website.dayOfAccess)
	fmt.Println("Month of access:")
	fmt.Scanf("%s", &website.monthOfAccess)
	fmt.Println("Year of access:")
	fmt.Scanf("%s", &website.yearOfAcess)
	return ReturnWebReference(&website)
}

//ReturnWebReference returns the information that user sent in a template rady for use in bibliography
func ReturnWebReference(w * websiteReference) string{
	var template string
	w.lastName= strings.ToUpper(w.lastName)
	template = fmt.Sprintf("%s, %s. %s. %s, %s. Disponível em: %s. Acesso em: %s %s. %s.", w.lastName, w.firstName, w.articleTitle, w.websiteName, w.year,  w.url, w.dayOfAccess, w.monthOfAccess, w.yearOfAcess)
	return template
}






