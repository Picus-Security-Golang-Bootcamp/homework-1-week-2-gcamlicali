package main

import (
	"fmt"
	"os"
	"strings"
)

// That slice contains booknames on the self as string
var BookListInShelf = []string{ 
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
}

func main() {

	var command string = strings.ToLower(os.Args[1]) // Converts command to lower string
	
	switch (command)	{ // Choose right command
	case "list":
		PrintBookList()
	case "search":
		SearchBookOnShelf(os.Args[2:])
	default:
		fmt.Println("Unknown Command! Please use 'list' or 'search <bookname>' commands") // if the command does not exist print this line
	}	
}


// Print All Books Name on New Line
func PrintBookList() {
	for i := 0; i < len(BookListInShelf); i++ {
		fmt.Println(BookListInShelf[i])
	}
}

// Search Book name in Booklist, print screen if it exists or not exist
func SearchBookOnShelf(booknameArr []string) {
	booknameStr := strings.ToLower(strings.Join(booknameArr, " ")) // convert to lower for right compare
	
	if (Contains(BookListInShelf, booknameStr)){ // is exist or not exist 
		fmt.Println("Yes, we have this book on our shelf")
	}	else{
		fmt.Println("Sorry, we don't have this book")
	}
}

// Compare string with other string in slice
func Contains(a []string, x string) bool {
    for _, n := range a {
        if x == strings.ToLower(n) {
            return true
        }
    }
    return false
}