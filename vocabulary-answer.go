package vocabulary4mydictionary

const (
	// Collection : int for "Type"
	Collection = 1
	// Dictionary : int for "Type"
	Dictionary = 2
	// Online : int for "Type"
	Online = 3
	// Basic : string for "Status"
	Basic = "basic"
	// Advance : string for "Status"
	Advance = "advance"
)

// VocabularyAnswerStruct : define and information
type VocabularyAnswerStruct struct {
	Word         string   `json:"word"`         // `xlsx:wd`
	Define       []string `json:"define"`       // `xlsx:def`
	SerialNumber int      `json:"serialNumber"` // `xlsx:sn`
	QueryCounter int      `json:"queryCounter"` // `xlsx:qc`
	QueryTime    string   `json:"queryTime"`    // `xlsx:qt`
	SourceName   string   `json:"sourceName"`
	Type         int      `json:"type"`
	Status       string   `json:"status"`
}
