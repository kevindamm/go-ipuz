package ipuz

import (
	"reflect"
	"testing"
)

func TestCrossword_MarshalJSON(t *testing.T) {
	type fields struct {
		Metadata    Metadata
		Dimensions  CrosswordDimensions
		Puzzle      [][]LabeledCell
		Solution    [][]CrosswordValue
		AcrossClues []Clue
		DownClues   []Clue
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
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
			got, err := cw.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Crossword.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Crossword.MarshalJSON() = %v, want %v", got, tt.want)
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