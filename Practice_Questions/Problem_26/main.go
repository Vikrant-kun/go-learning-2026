package main

import "fmt"

func main() {
	users := map[int]string{
		101: "Vikrant",
		102: "Ankit",
		103: "Rahul",
	}
	var id int

	fmt.Print("Enter User ID: ")
	fmt.Scan(&id)

	fmt.Println("Username: ", users[id])
}

