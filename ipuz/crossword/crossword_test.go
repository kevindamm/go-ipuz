package crossword

import (
	"bytes"
	"testing"
)

// TODO testing
// dimensions and reading Puzzle
// . LabeledCell as int
// . LabeledCell as string
// . LabeledCell as { cell: int, style: StyleSpec }
// dimensions and reading Solution
// . CrosswordValue as "#" block in solution
// . CrosswordValue as /[A-Z]{1}/ from string
// . CrosswordValue as {
// dimensions and reading Saved
// . CrosswordValue as "#" (must match block in Puzzle)
// . CrosswordValue as /[A-Z]{1} from string
// reading AcrossClues (decoded from clues["across"])
// reading DownClues (decoded from clues["down"])

func TestLabeledCell_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		fields  LabeledCell
		want    []byte
		wantErr bool
	}{
		{"zero value", LabeledCell{0, "", nil}, []byte("0"), false},
		{"labeled index", LabeledCell{Index: 42}, []byte("42"), false},
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var cw LabeledCell
			got, err := cw.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("LabeledCell.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !bytes.Equal(got, tt.want) {
				t.Errorf("LabeledCell.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCrossword_UnmarshalJSON(t *testing.T) {
	type fields struct {
		Metadata    Metadata
		Dimensions  CrosswordDimensions
		Puzzle      [][]LabeledCell
		Solution    [][]CrosswordValue
		AcrossClues []Clue
		DownClues   []Clue
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cw := &Crossword{
				Metadata:    tt.fields.Metadata,
				Dimensions:  tt.fields.Dimensions,
				Puzzle:      tt.fields.Puzzle,
				Solution:    tt.fields.Solution,
				AcrossClues: tt.fields.AcrossClues,
				DownClues:   tt.fields.DownClues,
			}
			if err := cw.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Crossword.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
