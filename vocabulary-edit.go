package vocabulary4mydictionary

const (
	// Define : int for "ContentType" in VocabularyEditStruct
	Define = -1
	// Note : int for "ContentType" in VocabularyEditStruct
	Note = -2
)

// VocabularyEditStruct : edit define or note of vocabulary
type VocabularyEditStruct struct {
	TableType int `json:"tableType"`
	Location  struct {
		TableIndex int `json:"tableIndex"`
		ItemIndex  int `json:"itemIndex"`
	} `json:"location"`
	ContentType int    `json:"contentType"`
	Content     string `json:"content"`
}
