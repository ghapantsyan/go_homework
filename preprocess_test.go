package main

import (
	"testing"
)

func TestTextPreprocess(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		ignoreFields int
		ignoreChars  int
		ignoreCase   bool
		want         []string
		wantErr      bool
	}{
		{
			name:         "basic processing",
			input:        "Hello\nWorld",
			ignoreFields: 0,
			ignoreChars:  0,
			ignoreCase:   false,
			want:         []string{"Hello", "World"},
			wantErr:      false,
		},
		{
			name:         "ignore case",
			input:        "HeLLo\nWoRLD",
			ignoreFields: 0,
			ignoreChars:  0,
			ignoreCase:   true,
			want:         []string{"hello", "world"},
			wantErr:      false,
		},
		{
			name:         "ignore fields within bounds",
			input:        "12345 67890\nabcde fghij",
			ignoreFields: 2,
			ignoreChars:  0,
			ignoreCase:   false,
			want:         []string{"345 67890", "cde fghij"},
			wantErr:      false,
		},
		{
			name:         "ignore chars within bounds",
			input:        "12345\n67890",
			ignoreFields: 0,
			ignoreChars:  2,
			ignoreCase:   false,
			want:         []string{"345", "890"},
			wantErr:      false,
		},
		{
			name:         "ignore fields out of range",
			input:        "short\nlines",
			ignoreFields: 10,
			ignoreChars:  0,
			ignoreCase:   false,
			want:         nil,
			wantErr:      true,
		},
		{
			name:         "ignore chars out of range",
			input:        "tiny\nline",
			ignoreFields: 0,
			ignoreChars:  10,
			ignoreCase:   false,
			want:         nil,
			wantErr:      true,
		},
		{
			name:         "combined processing",
			input:        "  HELLO  WORLD\n  FOO  BAR  ",
			ignoreFields: 1,
			ignoreChars:  2,
			ignoreCase:   true,
			want:         []string{"llo  world", "o  bar  "},
			wantErr:      false,
		},
		{
			name:         "empty input",
			input:        "",
			ignoreFields: 0,
			ignoreChars:  0,
			ignoreCase:   false,
			want:         []string{""},
			wantErr:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TextPreprocess(
				tt.input,
				tt.ignoreFields,
				tt.ignoreChars,
				tt.ignoreCase,
			)

			if (err != nil) != tt.wantErr {
				t.Errorf("TextPreprocess() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(got) != len(tt.want) {
				t.Errorf("TextPreprocess() = %v, want %v", got, tt.want)
				return
			}

			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("TextPreprocess()[%d] = %v, want %v", i, got[i], tt.want[i])
				}
			}
		})
	}
}
