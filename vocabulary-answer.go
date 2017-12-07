package vocabulary4mydictionary

const (
	// Collection : int for "Type" in VocabularyAnswerStruct and "TableType" in VocabularyEditStruct
	Collection = 1
	// Dictionary : int for "Type" in VocabularyAnswerStruct and "TableType" in VocabularyEditStruct
	Dictionary = 2
	// Online : int for "Type" in VocabularyAnswerStruct
	Online = 3
	// Basic : string for "Status" in VocabularyAnswerStruct
	Basic = "basic"
	// Advance : string for "Status" in VocabularyAnswerStruct
	Advance = "advance"
)

// VocabularyAnswerStruct : define and information
type VocabularyAnswerStruct struct {
	Word         string   `json:"word"`         // `xlsx:wd`
	Define       []string `json:"define"`       // `xlsx:def`
	SerialNumber int      `json:"serialNumber"` // `xlsx:sn`
	QueryCounter int      `json:"queryCounter"` // `xlsx:qc`
	QueryTime    string   `json:"queryTime"`    // `xlsx:qt`
	Note         []string `json:"note"`         // `xlsx:nt`
	SourceName   string   `json:"sourceName"`
	Type         int      `json:"type"`
	Status       string   `json:"status"`
	Location     struct {
		TableIndex int `json:"tableIndex"`
		ItemIndex  int `json:"itemIndex"`
	} `json:"location"`
}
