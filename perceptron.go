package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var numbers = []string{
	"zero", "one", "two", "three", "four",
	"five", "six", "seven", "eight", "nine"}

var endElementDigitInStringSlice = 900
var numberOfIterations = 5000000//9999999//
var currentNumber int
var numberOfPixel = 784
var threshold = 5000
var weights [784]int
var digitInStringSlice [][]string

func getDigitInStringSlice(fileName string) []string {
	csvFile, _ := os.Open(fileName)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	line, _ := reader.Read()
	csvFile.Close()
	return line[:endElementDigitInStringSlice]
}

func education() {
	var digit int
	var index int
	var flag bool
	fmt.Println("currentNumber = " + strconv.Itoa(currentNumber))
	for i := 0; i < numberOfIterations; i++ {
		digit = rand.Intn(10)
		index = rand.Intn(endElementDigitInStringSlice)
		flag = discrimination(digit, index)
		if digit == currentNumber {
			if !flag {
				editweights(digitInStringSlice[digit][index], 20)
			}
		} else {
			if flag {
				editweights(digitInStringSlice[digit][index], -1)
			}
		}
	}
	return
}

var o int

func discrimination(digit int, index int) bool {
	sum := 0
	digit = 9
	for i := 0; i < numberOfPixel; i++ {
		//fmt.Println("i = " + strconv.Itoa(i))
		if digitInStringSlice[digit][index][i] == '1' {
			sum += weights[i]
		}
	}
	// fmt.Println(sum)
	// if o > 5 {
	// 	os.Exit(0)
	// }
	// o++
	return sum >= threshold
}

func editweights(digit string, cof int) {
	for i := 0; i < numberOfPixel; i++ {
		if digit[i] == '1' {
			weights[i] += cof
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	currentNumber, _ = strconv.Atoi(os.Args[1])
	digitInStringSlice = make([][]string, 10)
	//d := 0
	//k := getDigitInStringSlice("CSVs" + string(os.PathSeparator) + strconv.Itoa(1) + ".csv")
	//for i := 0; i < 10; i++ {
	//	l := rand.Intn(800)
	//	p := 0
	//	for j := 0; j < 784; j++ {
	//		if(k[l][j] == '1') {
	//			p++
	//		}
	//	}
	//	d += p
	//	fmt.Print(strconv.Itoa(i) + ": " )
	//	fmt.Print("l = " + strconv.Itoa(l))
	//	fmt.Println(";\tp = " + strconv.Itoa(p))
	//}
	//fmt.Println("d/10 = " + strconv.Itoa(d/10))
	//os.Exit(0)
	for i := 0; i < 10; i++ {
		digitInStringSlice[i] = getDigitInStringSlice("CSVs" + string(os.PathSeparator) + strconv.Itoa(i) + ".csv")
	}
	education()

	index := 0
	for i := 0; i < 10; i++ {
		index = rand.Intn(100) + 800
		fmt.Print(i)
		//fmt.Print("\t")
		//fmt.Print(index)
		fmt.Print("\t")
		if discrimination(i, index) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}

	// fmt.Println()
	// for index := 0; index < 784; index++ {
	// 	if weights[index] <= 0{
	// 		fmt.Println(weights[index])
	// 	}
	// }
	//fmt.Println()
	//for index := 0; index < 784; index++ {
	//	if weights[index] > 0{
	//		fmt.Println(weights[index])
	//	}
	//}
}
