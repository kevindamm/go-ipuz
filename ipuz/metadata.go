package ipuz

// These metadata entries are considered common among the other ipuz variants.
//
// The fields are separated into related groups -- provenance,
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

