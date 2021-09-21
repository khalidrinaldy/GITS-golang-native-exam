package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var menu int
var ID int = 0
var title, singer, saver string
var vote float64
var data = [][4]string{{"ID", "Title", "Singer", "Vote"}}

func add() {
	fmt.Println("Input singer title : ")
	fmt.Scan(&title)
	ID++
	fmt.Println("Input singer : ")
	fmt.Scan(&singer)
	fmt.Println("Input your vote : ")
	fmt.Scan(&vote)
	if valid, err := validate(singer); valid {
		newdata := [4]string{(strconv.Itoa(ID)), title, singer, (strconv.FormatFloat(vote, 'f', 1, 64))}
		data = append(data, newdata)
		newdata = [4]string{}
	} else {
		fmt.Println(err.Error())
	}
}
func showAll() {
	fmt.Println("List of singer available : \n", len(data))
	for i := 0; i < len(data); i++ {
		for j := 0; j < 4; j++ {
			fmt.Print(data[i][j], "  \t")
		}
		fmt.Println("")
	}
}

func searchID(ID int) int {
	var index int
	search := strconv.Itoa(ID)
	for i := 0; i < len(data); i++ {
		if data[i][0] == search {
			index = i
			break
		}
	}
	return index
}

func deleteID(ID int) {
	var before, after [][4]string
	index := searchID(ID)
	before = data[:index]
	after = data[index+1:]
	data = [][4]string{}
	for i := 0; i < len(before); i++ {
		data = append(data, before[i])
	}
	for i := 0; i < len(after); i++ {
		data = append(data, after[i])
	}
}

func searchSinger() {
	fmt.Println("Search result(s) : ")
	for i := 0; i < len(data); i++ {
		matched, _ := regexp.MatchString(`^a|^A`, data[i][2])
		if matched == true {
			fmt.Println("ID : ", data[i][0])
			fmt.Println("Title : ", data[i][1])
			fmt.Println("Singer : ", data[i][2])
			fmt.Println("Vote : ", data[i][3])
		}
	}
}
func stringtoFloat(input string) float64 {
	var save float64
	if s, err := strconv.ParseFloat(input, 64); err == nil {
		save = s
	}
	return save
}
func topthree() {
	if len(data) < 4 {
		fmt.Println("Sorry the data is too short, please add more")
	} else {
		var third, first, second [4]string
		for i := 1; i < len(data); i++ {
			if stringtoFloat(data[i][3]) > stringtoFloat(first[3]) {
				third = second
				second = first
				first = data[i]
			} else if stringtoFloat(data[i][3]) > stringtoFloat(second[3]) {
				third = second
				second = data[i]
			} else if stringtoFloat(data[i][3]) > stringtoFloat(third[3]) {
				third = data[i]
			}
		}
		fmt.Print("Top 3 most voted singer : \n")
		fmt.Print(data[0][0], "\t", data[0][1], "\t", data[0][2], "\t", data[0][3], "\n")
		fmt.Print(first[0], "\t", first[1], "\t", first[2], "\t", first[3], "\n")
		fmt.Print(second[0], "\t", second[1], "\t", second[2], "\t", second[3], "\n")
		fmt.Print(third[0], "\t", third[1], "\t", third[2], "\t", third[3], "\n")
	}
}
func countVote() float64 {
	var count float64 = 0
	for i := 1; i < len(data); i++ {
		count = count + stringtoFloat(data[i][3])
		fmt.Println(count)
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
		var num, err = strconv.Atoi(saver)
		if err == nil && num > 0 {
			deleteID(num)
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
