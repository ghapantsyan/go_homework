package main

import (
	"flag"
	"fmt"
)

func main() {

	const TEXT = `I love music.
I love music.
I love music.

I love music of Kartik.
I love music of Kartik.
Thanks.
I love music of Kartik.
I love music of Kartik.`

	uniqs := flag.Bool("u", false, "print uniqs only")
	doubles := flag.Bool("d", false, "print doubles only")
	counts := flag.Bool("c", false, "print row counts")
	ignoreCase := flag.Bool("i", false, "ignore case of text")
	ignoreFields := flag.Int("f", 0, "ignore first num fields in row")
	ignoreChars := flag.Int("s", 0, "ignore first num chars in row")

	flag.Parse()

	typeOfPrint := GetTypeOfPrint(*uniqs, *doubles, *counts)

	processedText, err := TextPreprocess(TEXT, *ignoreFields, *ignoreChars, *ignoreCase)
	if err != nil {
		fmt.Println(err)
		fmt.Println("ignore indexes will be skiped")
		processedText, _ = TextPreprocess(TEXT, 0, 0, *ignoreCase)
	}

	rowCounts := GetRowCounts(processedText)

	ResultPrinter(rowCounts, TEXT, typeOfPrint)
}
