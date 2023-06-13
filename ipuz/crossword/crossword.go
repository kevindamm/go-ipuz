package crossword

import (
	"io"

	"github.com/kevindamm/ipuz"
)

func LoadCrossword(fileData io.Reader) (Crossword, error) {
	var cwData Crossword


	return cwData, nil
}

type Crossword struct {
	ipuz.Metadata

	Dimensions  CrosswordDimensions `json:"dimensions"`
	Puzzle      [][]LabeledCell     `json:"puzzle"`
	Saved       [][]CrosswordValue  `json:"saved,omitempty"`
	Solution    [][]CrosswordValue  `json:"solution"`
	Clues       CrosswordClues      `json:"clues"`
}

type CrosswordDimensions struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type LabeledCell interface {
	Index() int  // if a clue starts in this cell
	Value() rune // BLOCK, NULL or pre-filled letter
	Style() *StyleSpec
}

type crosswordCell struct {
	Value int32
	Style *StyleSpec
}

func (cell *crosswordCell) MarshalJSON() ([]byte, error) {
  if cell.Index != 0 {
		if cell.Value != 0 || cell.Style != nil {
			return marshalLabeledCellAsDict(cell)
		}
	}
}

func marshalLabeledCellAsDict(cell *crosswordCell)

const (
	BLOCK_VALUE rune = '#'
	NULL_VALUE  rune = 0x7F
)

// Common zero value, serializes as int value 0.
var ZERO_CELL LabeledCell = LabeledCell{}

type CrosswordValue struct {
	Value rune
	Style StyleSpec
}

type CrosswordClues struct {
	Across []clue `json:"across"`
	Down   []clue `json:"down"`
}
	
type clue struct {
	Index int
	Text  string
}