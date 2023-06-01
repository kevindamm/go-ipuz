package ipuz

type Metadata struct {
	Version string   `json:"version"`
	Kind    []string `json:"kind"`

	Copyright   string `json:"copyright,omitempty"`
	Publisher   string `json:"publisher,omitempty"`
	Publication string `json:"publication,omitempty"`
	URL         string `json:"url,omitempty"`
	UniqueID    string `json:"uniqueid,omitempty"`

	Title       string `json:"title,omitempty"`
	Intro       string `json:"intro,omitempty"`
	Explanation string `json:"explanation,omitempty"`
	Annotation  string `json:"annotation,omitempty"`

	Author     string `json:"author,omitempty"`
	Editor     string `json:"editor,omitempty"`
	Date       string `json:"date,omitempty"`
	Origin     string `json:"origin,omitempty"`
	Notes      string `json:"notes,omitempty"`
	Difficulty string `json:"difficulty,omitempty"`
	Charset    string `json:"charset,omitempty"`

	Block string `json:"block,omitempty"` // defaults to #
	Empty string `json:"empty,omitempty"` // defaults to 0

	Check  SolutionChecksum     `json:"checksum,omitempty"`
	Styles map[string]StyleSpec `json:"styles,omitempty"`

	// a "saved" field is also common but takes a different shape for each puzzle
	// so it is left to the puzzle-specific struct definitions where the encoding
	// and decoding operations are also defined.
}

type SolutionChecksum []string

type StyleSpec struct {
	ShapeBackground string `json:"shapebg,omitempty"` // typically circle
	Highlight       bool   `json:"highlight,omitempty"`
	Named           string `json:"named,omitempty"`   // specifies delimiter for group name
	Border          int    `json:"border,omitempty"`  // border thickness
	Divided         string `json:"divided,omitempty"` // may be '-' '|' '/' '\' '+' 'x'

	Label           string            `json:"label,omitempty"`
	Mark            map[string]string `json:"mark,omitempty"`    // key is TL|TR|BL|BR|T|L|R|B|C
	ImageBackground string            `json:"imagebg,omitempty"` // URL
	Image           string            `json:"image,omitempty"`   // URL
	Slice           []string          `json:"slice,omitempty"`   // col1, row1, col2, row2

	Barred      string `json:"barred,omitempty"`      // TRBL
	Dotted      string `json:"dotted,omitempty"`      // TRBL
	Dashed      string `json:"dashed,omitempty"`      // TRBL
	LessThan    string `json:"lessthan,omitempty"`    // TRBL
	GreaterThan string `json:"greaterthan,omitempty"` // TRBL
	Equal       string `json:"equal,omitempty"`       // TRBL

	// Following fields have a value that is either a number (index into the color table) or a
	// hexadecimal string in the format of "hhhhhh".
	Color       ColorValue `json:"color,omitempty"`
	ColorText   ColorValue `json:"colortext,omitempty"`
	ColorBorder ColorValue `json:"colorborder,omitempty"`
	ColorBar    ColorValue `json:"colorbar,omitempty"`
}

type ColorValue struct {
	Index int    // will be -1 if undefined, or 0 to at least 16 for indices into the color table
	Value string // "hhhhhh" value for RGB hexadecimals.  May be undefined if Index is defined.
}

// TODO: MarshalJSON and UnmarshalJSON implementations for ColorValue

type GroupSpec struct {
	Rectangle []int        `json:"rect,omitempty"` // [ col1, row1, col2, row2 ]
	Line      interface{}  `json:"line,omitempty"` // [ [col1, row1], [deltacol, deltarow], length ]
	Cells     []Coordinate `json:"cells,omitempty"`
	Style     StyleSpec    `json:"style,omitempty"`
}

type Coordinate []int // Typically 2-dimensional
