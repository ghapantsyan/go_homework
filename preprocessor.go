package main

import (
	"errors"
	"fmt"
	"strings"
)

func TextPreprocess(
	text string,
	ignoreFields int,
	ignoreChars int,
	ignoreCase bool,
) (slicedText []string, err error) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			err = errors.New("ignore index out of range")
			slicedText = nil
		}
	}()

	if ignoreCase {
		text = strings.ToLower(text)
	}
	slicedText = strings.Split(text, "\n")

	if ignoreFields > 0 {
		for index := range slicedText {
			slice := strings.Split(slicedText[index][ignoreFields:], " ")
			if ignoreFields >= len(slice) {
				panic("ignoreFields too big")
			}
			newString := strings.Join(slice, " ")
			slicedText[index] = newString
		}
	}

	if ignoreChars > 0 {
		for index := range slicedText {
			slice := strings.Split(slicedText[index][ignoreChars:], "")
			if ignoreChars >= len(slice) {
				panic("ignoreChars too big")
			}
			newString := strings.Join(slice, "")
			slicedText[index] = newString
		}
	}

	return
}
