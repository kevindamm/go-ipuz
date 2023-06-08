package ipuz

type Crossword struct {
	Metadata

	Dimensions  CrosswordDimensions `json:"dimensions"`
	Puzzle      [][]LabeledCell     `json:"puzzle"`
	Solution    [][]CrosswordValue  `json:"solution"`
	AcrossClues []Clue
	DownClues   []Clue
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