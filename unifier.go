package main

import (
	"strings"
)

func GetRowCounts(slicedText []string) (rowCnts map[int]int) {

	rowCnts = make(map[int]int, len(slicedText))

	for index, value := range slicedText {
		if index == 0 {
			rowCnts[index] = 1
		} else if slicedText[index-1] != value {
			rowCnts[index] = 1
		} else {
			rowCnts[index] = 1 + rowCnts[index-1]
			rowCnts[index-1] = -1
		}
	}
	return
}

func GetRowCounts2(initText string) (counts map[string]int) { // depricated

	counts = make(map[string]int)

	newSlice := strings.Split(initText, "\n")

	for _, value := range newSlice {
		if counts[value] != 0 {
			counts[value]++
		} else {
			counts[value] = 1
		}
	}
	return
}
