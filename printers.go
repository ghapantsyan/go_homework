package main

import (
	"fmt"
	"strings"
)

func ResultPrinter(rowCnts map[int]int, originalText string, typeOfPrint string) {

	slicedText := strings.Split(originalText, "\n")
	var condition func(cnt int) bool
	printCnt := false

	if typeOfPrint == "doubles" {
		condition = func(cnt int) bool {
			return cnt > 1
		}
	} else if typeOfPrint == "uniqs" {
		condition = func(cnt int) bool {
			return cnt == 1
		}
	} else if typeOfPrint == "counts" {
		condition = func(cnt int) bool {
			return cnt > 0
		}
		printCnt = true
	} else {
		condition = func(cnt int) bool {
			return cnt > 0
		}
	}

	for index := range slicedText {
		if condition(rowCnts[index]) {
			if !printCnt {
				fmt.Println(slicedText[index])
			} else {
				fmt.Println(rowCnts[index], slicedText[index])
			}
		}
	}
}

func GetTypeOfPrint(uniqs bool, doubles bool, counts bool) string {

	var counter int8
	typeOfPrint := ""

	if uniqs {
		typeOfPrint = "uniqs"
		counter++
	}
	if doubles {
		typeOfPrint = "doubles"
		counter++
	}
	if counts {
		typeOfPrint = "counts"
		counter++
	}
	if counter == 1 {
		return typeOfPrint
	} else {
		return ""
	}
}
