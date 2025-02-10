package main

import (
	"strings"
)

func GetRowCounts(slicedText []string) (row_cnts map[int]int) {

	row_cnts = make(map[int]int, len(slicedText))

	for index, value := range slicedText {
		if index == 0 {
			row_cnts[index] = 1
		} else if slicedText[index-1] != value {
			row_cnts[index] = 1
		} else {
			row_cnts[index] = 1 + row_cnts[index-1]
			row_cnts[index-1] = -1
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
