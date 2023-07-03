package main

import(
	"os"
	"fmt"
)


func usage(){
	fmt.Println("You need to use at least 1 argument\n-b for books\n-w for websites\nIf you found any bugs e-mail me")
}

func main(){
	userArg := os.Args[1]

	if userArg == "-b" {
		fmt.Println(GetBookInfo())
	}else if userArg == "-w"{
		fmt.Println(GetWebsiteInfo())
	}else{
		usage()
	}
}
