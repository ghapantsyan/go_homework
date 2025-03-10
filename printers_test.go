package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestGetTypeOfPrint(t *testing.T) {
	tests := []struct {
		uniqs   bool
		doubles bool
		counts  bool
		want    string
	}{
		{uniqs: false, doubles: false, counts: false, want: ""},
		{uniqs: true, doubles: false, counts: false, want: uniqs_type},
		{uniqs: false, doubles: true, counts: false, want: doubles_type},
		{uniqs: false, doubles: false, counts: true, want: counts_type},
		{uniqs: true, doubles: true, counts: false, want: ""},
		{uniqs: false, doubles: true, counts: true, want: ""},
		{uniqs: true, doubles: false, counts: true, want: ""},
		{uniqs: true, doubles: true, counts: true, want: ""},
	}

	for _, tt := range tests {
		got := GetTypeOfPrint(tt.uniqs, tt.doubles, tt.counts)
		if got != tt.want {
			t.Errorf("GetTypeOfPrint(%v, %v, %v) = %v, want %v",
				tt.uniqs, tt.doubles, tt.counts, got, tt.want)
		}
	}
}

func TestResultPrinter(t *testing.T) {
	tests := []struct {
		name         string
		rowCnts      map[int]int
		originalText string
		printType    string
		wantOutput   string
	}{
		{
			name:         "doubles",
			rowCnts:      map[int]int{0: 2, 1: 1, 2: 3},
			originalText: "a\nb\nc",
			printType:    doubles_type,
			wantOutput:   "a\nc\n",
		},
		{
			name:         "uniqs",
			rowCnts:      map[int]int{0: 1, 1: 2, 2: 1},
			originalText: "x\ny\nz",
			printType:    uniqs_type,
			wantOutput:   "x\nz\n",
		},
		{
			name:         "counts",
			rowCnts:      map[int]int{0: 5, 1: 0, 2: 3},
			originalText: "line1\nline2\nline3",
			printType:    counts_type,
			wantOutput:   "5 line1\n3 line3\n",
		},
		{
			name:         "invalid_type",
			rowCnts:      map[int]int{0: 1, 1: 0},
			originalText: "a\nb",
			printType:    "invalid",
			wantOutput:   "a\n",
		},
		{
			name:         "empty_lines",
			rowCnts:      map[int]int{0: 3, 1: 0, 2: 2},
			originalText: "\n\n\n",
			printType:    doubles_type,
			wantOutput:   "\n\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Redirect stdout
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			ResultPrinter(tt.rowCnts, tt.originalText, tt.printType)

			w.Close()
			os.Stdout = old

			var buf bytes.Buffer
			io.Copy(&buf, r)

			if got := buf.String(); got != tt.wantOutput {
				t.Errorf("Unexpected output:\nGot: %q\nWant: %q", got, tt.wantOutput)
			}
		})
	}
}
