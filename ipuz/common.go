package ipuz

type StyleSpec struct {
	ShapeBackground string `json:"shapebg,omitempty"` // typically circle
	Highlight       bool   `json:"highlight,omitempty"`
	Named           string `json:"named,omitempty"`   // specifies delimiter for group name
	Border          int    `json:"border,omitempty"`  // border thickness
	Divided         string `json:"divided,omitempty"` // may be '-' '|' '/' '\' '+' 'x'

	Label           string `json:"label,omitempty"`
	ImageBackground string `json:"imagebg,omitempty"` // URL
	Image           string `json:"image,omitempty"`   // URL

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
