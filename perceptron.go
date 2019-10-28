package main

import (
	"os"
	"fmt"
	"strconv"
	"math/rand"
	"encoding/csv"
	"bufio"
)

var numbers = []string{
	"zero", "one", "two", "three", "four",
	"five", "six", "seven", "eight", "nine"}

var endElementDigitInStringSlice = 900
var numberOfIterations = 1000000;
var currentNumber int;
var numberOfPixel = 784;
var threshold = 300;
var weights [784]int;
var digitInStringSlice [][]string

func getDigitInStringSlice(fileName string) []string {
	csvFile, _ := os.Open(fileName)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	line, _ := reader.Read()
	for i := 0; i < 800; i++ {
		return line
	}
	csvFile.Close()
	return line
}

func education() {
	var digit int
	var index int
	var flag bool
	for i := 0; i < numberOfIterations; i++ {
		digit = rand.Intn(10)
		index = rand.Intn(endElementDigitInStringSlice)
		flag = proceed(digit, index)
		if digit == currentNumber {
			if (!flag) {
				editweights(digitInStringSlice[digit][index], -1)
			}
		} else {
			if (flag) {
				editweights(digitInStringSlice[digit][index], 1)
			}
		}
	}
	return
}

func proceed(digit int, index int) bool {
	sum := 0
	for i := 0; i < numberOfPixel; i++{
		if digitInStringSlice[digit][index][i] == '1'{
			sum += weights[i]
		}
	}
	return sum >= threshold
}

func editweights(digit string, cof int)  {
	for i := 0; i < numberOfPixel; i++ {
		if(int(digit[i]) == 1) {
			weights[i] +=  cof
		}
	}
}

func main() {
	currentNumber, _ := strconv.Atoi(os.Args[1])
	digitInStringSlice = make([][]string, 10)
	for i := 0; i < 10; i++ {
		digitInStringSlice[i] = getDigitInStringSlice("CSVs" + string(os.PathSeparator) + strconv.Itoa(i) + ".csv")
	}
	education()

	index := 0
	for i := 0; i < 10; i++ {
		index = rand.Intn(100) + 800
		if proceed(currentNumber, index) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
	for index := 0; index < 784; index++ {
		fmt.Println(weights[index])
	}
}
