package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

// That slice contains booknames on the self as string
var BookShelf = []string{ 
	"Anna Karenina", 
	"To Kill a Mockingbird",
	"The Great Gatsby",
	"One Hundred Years of Solitude",
	"A Passage to India",
	"Invisible Man",
	"Don Quixote",
	"Beloved",
	"Mrs. Dalloway",
	"Things Fall Apart",
	"Jane Eyre",
	"The Color Purple",
	"Early in Orcadia",
	"Raven Black",
	"Invisible Islands",
	"The Blackhouse",
	"Why the Whales Came",
	"Where is Anna",
	"The Man far from Earth",
	"Black Sails"}

func main() {

	command:= strings.ToLower(os.Args[1]) // Converts command to lower string
	inputBookName := os.Args[2:]

	switch (command)	{ // Choose right command
	case "list":
		PrintBookList()
	case "search":
		SearchBookOnShelf(inputBookName)
	default:
		fmt.Println("Unknown Command! Please use 'list' or 'search <bookname>' commands") // if the command does not exist print this line
	}	
}


// Print All Books Name on New Line
func PrintBookList() {
	for i := 0; i < len(BookShelf); i++ {
		fmt.Println(BookShelf[i])
	}
}

// Search Book name in Booklist, print screen if it exists or not exist
func SearchBookOnShelf(booknameArr []string) {
	
	var BookId []int
	booknameStr := strings.ToLower(strings.Join(booknameArr, " ")) // convert input to lower for right compare
	
	for i := 0; i < len(BookShelf); i++ {
		if strings.Contains(strings.ToLower(BookShelf[i]), booknameStr) {
			BookId = append(BookId,i)
		}
	}

	if (len(BookId) > 0) {

		fmt.Print("We found ")
		fmt.Print(strconv.Itoa((len(BookId))))  // convert int to str
		fmt.Println(" book(s) on shelf : ")

		for i := 0; i < len(BookId); i++ {
			fmt.Println(BookShelf[BookId[i]]) // Prints found books from list
		}
	} else {
		fmt.Println("Sorry, we don't have this book")
	}
}
