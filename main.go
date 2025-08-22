package main

import "fmt"

func main(){

	var from string
	fmt.Println("Write from:")
	fmt.Scanf("%s", &from)

	var to string;
	fmt.Println("Write to:")
	fmt.Scanf("%s",&to)

	var passwordApp string
	fmt.Println("Write app password:")
	fmt.Scanf("%s", &passwordApp);


	fmt.Printf("%s %s %s", from, to, passwordApp);
}