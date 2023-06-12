package ipuz

import (
	"encoding/json"
)

type Crossword struct {
	Metadata

	Dimensions  CrosswordDimensions `json:"dimensions"`
	Puzzle      [][]LabeledCell     `json:"puzzle"`
	Saved       [][]CrosswordValue  `json:"saved"`
	Solution    [][]CrosswordValue  `json:"solution"`
	AcrossClues []Clue
	DownClues   []Clue
}

func (cw *Crossword) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Dimensions CrosswordDimensions`json:"dimensions"`
		Puzzle      [][]LabeledCell     `json:"puzzle"`
		Solution    [][]CrosswordValue  `json:"solution"`
	}{
		Dimensions:       cw.Dimensions,
		Puzzle: [][]LabeledCell{},
		Solution: [][]CrosswordValue{},
	})
}

func (cw *Crossword) UnmarshalJSON(data []byte) error {
	type Alias Crossword
	aux := &struct {
		AcrossClues []Clue
		DownClues []Clue
		*Alias
	}{
		Alias: (*Alias)(cw),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	cw.AcrossClues = aux.AcrossClues
	cw.DownClues = aux.DownClues
	return nil
}

type CrosswordDimensions struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type LabeledCell struct {
	Index int  // if a clue starts in this cell
	Value rune // BLOCK, NULL or pre-filled letter
	Style StyleSpec
}

const (
	BLOCK_VALUE rune = '#'
	NULL_VALUE  rune = 0x7F
)

type CrosswordValue struct {
	Value rune
	Style StyleSpec
}

type Clue struct {
	Index int
	Text  string
}