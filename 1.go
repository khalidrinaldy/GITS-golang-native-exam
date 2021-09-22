package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
)

var menu, saver, vote int
var ID int = 0
var title, singer string

type database struct {
	ID     int
	Title  string
	Singer string
	vote   int
}

var data []database

func add() {
	fmt.Println("Input singer's album : ")
	fmt.Scan(&title)
	ID++
	fmt.Println("Input singer : ")
	fmt.Scan(&singer)
	fmt.Println("Input your vote : ")
	fmt.Scan(&vote)
	if valid, err := validate(singer); valid {
		newdata := database{ID: ID, Title: title, Singer: singer, vote: vote}
		data = append(data, newdata)
		newdata = database{}
	} else {
		ID--
		fmt.Println(err.Error())
	}
}
func showAll() {
	fmt.Println("List of singer available : ", len(data))
	fmt.Println("ID", "\t\t", "Title", "\t\t", "Singer", "\t\t", "Vote")
	fmt.Println("-----------------------------------------------------")
	for _, list := range data {
		fmt.Println(list.ID, "\t\t", list.Title, "\t\t", list.Singer, "\t\t", list.vote)
		fmt.Println("-----------------------------------------------------")
	}
}

func searchID(ID int) int {
	var index int
	search := ID
	for i, list := range data {
		if list.ID == search {
			index = i
			break
		}
	}
	return index
}

func deleteID(ID int) {
	var before, after []database
	index := searchID(ID)
	before = data[:index]
	after = data[index+1:]
	data = data[:0]
	for _, listbef := range before {
		data = append(data, listbef)
	}
	for _, listaft := range after {
		data = append(data, listaft)
	}
}
func searchSinger() {
	fmt.Println("Search result(s) : ")
	for _, list := range data {
		matched, _ := regexp.MatchString(`^a|^A`, list.Singer)
		if matched == true {
			fmt.Println("--------------")
			fmt.Println("ID : ", list.ID)
			fmt.Println("Title : ", list.Title)
			fmt.Println("Singer : ", list.Singer)
			fmt.Println("Vote : ", list.vote)
			fmt.Println("--------------")
		}
	}
}

func topthree() {
	if len(data) < 3 {
		fmt.Println("Sorry the data is too short, please add more")
	} else {
		var third, first, second database
		for i, list := range data {
			if list.vote > first.vote {
				third = second
				second = first
				first = data[i]
			} else if list.vote > second.vote {
				third = second
				second = data[i]
			} else if list.vote > third.vote {
				third = data[i]
			}
		}
		fmt.Print("Top 3 most voted singer : \n")
		fmt.Print("ID", "\t\t", "Title", "\t\t", "Singer", "\t\t", "Vote", "\n")
		fmt.Println("-----------------------------------------------------")
		fmt.Print(first.ID, "\t\t", first.Title, "\t\t", first.Singer, "\t\t", first.vote, "\n")
		fmt.Println("-----------------------------------------------------")
		fmt.Print(second.ID, "\t\t", second.Title, "\t\t", second.Singer, "\t\t", second.vote, "\n")
		fmt.Println("-----------------------------------------------------")
		fmt.Print(third.ID, "\t\t", third.Title, "\t\t", third.Singer, "\t\t", third.vote, "\n")
	}
}
func countVote() int {
	var count int = 0
	for _, list := range data {
		count = count + list.vote
	}
	return count
}
func validate(input string) (bool, error) {
	if input == "" {
		return false, errors.New("cannot be empty")
	}
	m := regexp.MustCompile("[a-zA-Z]")
	if m.MatchString(input) == false {
		return false, errors.New("please input strings")
	}
	return true, nil
}

func main() {
	fmt.Println("Welcome to singer Review! Please select menu :\n1. Add new singer data\n2. Delete a singer data\n3. Show all singer data\n4. Count votes i\n5. Top 3 rated singers\n6. Singer whose name starts with 'A'\n7. Exit")
	fmt.Scan(&menu)
	if menu == 1 {
		add()
		main()
	} else if menu == 2 {
		fmt.Println("What ID you want to delete? : ")
		fmt.Scan(&saver)
		if saver > 0 {
			deleteID(saver)
		} else {
			fmt.Println("Sorry, please input again")
		}
		showAll()
		main()
	} else if menu == 3 {
		showAll()
		main()
	} else if menu == 4 {
		fmt.Println("Sum of all counted voutes : ", countVote())
		main()
	} else if menu == 5 {
		topthree()
		main()
	} else if menu == 6 {
		searchSinger()
		main()
	} else if menu == 7 {
		fmt.Println("Thanks")
		os.Exit(1)
	}

}
