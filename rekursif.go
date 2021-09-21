package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/lynn9388/supsub"
)

func main() {
	defer catch()
	var batasRekur1, batasRekur2 string
	var word string

	fmt.Print("\033[H\033[2J")
	fmt.Println("Rekursif 1")
	fmt.Print("Masukan batas rekursif 1 : ")
	fmt.Scan(&batasRekur1)
	if valid, err := validate(batasRekur1); !valid {
		panic(err.Error())
	}
	var amount1, _ = strconv.Atoi(batasRekur1)
	rekursif1(amount1, 5, 0)

	fmt.Println("\nRekursif 2")
	fmt.Print("Masukan batas rekursif 2 : ")
	fmt.Scan(&batasRekur2)
	if valid, err := validate(batasRekur2); !valid {
		panic(err.Error())
	}
	var amount2, _ = strconv.Atoi(batasRekur2)
	var param2 = Param2{amount: amount2, denominator: 1, finalDenominator: 1, sliceDenominator: []int{}}
	rekursif2(param2)

	fmt.Println("\nRekursif 3")
	fmt.Print("Masukan kata : ")
	fmt.Scan(&word)
	fmt.Println("Total huruf kapital :", rekursif3(word, 0, 0))
}

func rekursif1(amount, startNumber, total int)  {
	total+=startNumber
	if amount == 1 {
		fmt.Printf("%d = %d", startNumber, total)
	} else {
		fmt.Printf("%d + ", startNumber)
		rekursif1(amount-1, startNumber*10, total)
	}
}

func rekursif2(param2 Param2)  {
	param2.sliceDenominator = append(param2.sliceDenominator, int(math.Pow(float64(param2.denominator),2)))
	if param2.amount == 1 {
		fmt.Printf("x%v/%.0f = ", supsub.ToSup(strconv.Itoa(param2.denominator)), math.Pow(float64(param2.denominator),2))
		fmt.Print("(")
		for i := 0; i < param2.denominator; i++ {
			if i == param2.denominator-1 {
				fmt.Printf("%dx%v", param2.finalDenominator/param2.sliceDenominator[i],supsub.ToSup(strconv.Itoa(i+1)))
			} else {
				fmt.Printf("%dx%v + ", param2.finalDenominator/param2.sliceDenominator[i],supsub.ToSup(strconv.Itoa(i+1)))
			}
		}
		fmt.Printf(")/%d", param2.finalDenominator)
	} else {
		fmt.Printf("x%v/%.0f + ", supsub.ToSup(strconv.Itoa(param2.denominator)), math.Pow(float64(param2.denominator),2))
		param2.amount-=1
		param2.denominator+=1
		param2.finalDenominator*=int(math.Pow(float64(param2.denominator),2))
		rekursif2(param2)
	}
}

func rekursif3(word string, index, total int) int {
	sliceWord := strings.Split(word, "")
	if index == len(sliceWord) {
		return total
	}
	
	if []rune(sliceWord[index])[0] > 64 && []rune(sliceWord[index])[0] < 91 {
		total+=1
	}
	return rekursif3(word, index+1, total)
}

type Param2 struct {
	amount, denominator, finalDenominator int
	sliceDenominator []int
}

func validate(data string) (bool, error) {
	if data == "0" {
		return false, errors.New("tidak boleh 0")
	}
	
	_, err := strconv.ParseFloat(data, 64)
	if err == nil {
		return true, nil
	}
	return false, errors.New("input harus berupa nomor")
}

func catch() {
	if r := recover(); r != nil {
		fmt.Println("Error(\"", r, "\")")
	}
}