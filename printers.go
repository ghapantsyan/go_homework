package main

import (
	"fmt"
	"strings"
)

const doubles_type = "doubles"
const uniqs_type = "uniqs"
const counts_type = "counts"

func ResultPrinter(rowCnts map[int]int, originalText string, typeOfPrint string) {

	slicedText := strings.Split(originalText, "\n")
	var condition func(cnt int) bool
	printCnt := false

	if typeOfPrint == doubles_type {
		condition = func(cnt int) bool {
			return cnt > 1
		}
	} else if typeOfPrint == uniqs_type {
		condition = func(cnt int) bool {
			return cnt == 1
		}
	} else if typeOfPrint == counts_type {
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
		typeOfPrint = uniqs_type
		counter++
	}
	if doubles {
		typeOfPrint = doubles_type
		counter++
	}
	if counts {
		typeOfPrint = counts_type
		counter++
	}
	if counter == 1 {
		return typeOfPrint
	} else {
		return ""
	}
}
