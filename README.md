# Vocabulary for MYDICTIONARY

### 1. Introduction

The repository defines some basic data structures of [MYDICTIONARY](https://github.com/zzc-tongji/mydictionary).

### 2. Query

Methods of query should be explained first before discussing data structures.

Take "animal.xlsx" as an example.

![animal](./README.picture/animal.png)

#### 2.1. Basic Query

If the word **is exactly the same as the content in cell at column "Word" in a line**, the line will be selected.

In basic query, input word "cat" to file "animal.xlsx", and then, the line 1 will be selected.

#### 2.2. Advanced Query

If the word **is contained by the content in cells at columns "Word", "Define" or "Note" in a line**, the line will be selected.

In advanced query, input word "e" to "animal.xlsx", and then, the line 1, 3 and 5 will be selected.

### 3. Ask

```go
type VocabularyAskStruct struct {
	Word        string `json:"word"`
	Advance     bool   `json:"advance"`
	Online      bool   `json:"online"`
	DoNotRecord bool   `json:"doNotRecord"`
}
```

This structure indicates the word and the method for query.

#### 3.1. Advance

If `Advance` is `false`, the library will do *basic query* of `Word` only.

If `Advance` is `true`, the library will do both *basic query* and *advanced query* of `Word`.

#### 3.2. Online

If `Online` is `true`, the library will know that users need query `Word` online.

This doesn't ensure that the library will query `Word` online. Whether the library does also depends on `online.mode` in *configuration* (at [here](https://github.com/zzc-tongji/mydictionary#2431-mode)).

#### 3.3. DoNotRecord

If `DoNotRecord` is `true`, the library will not *record* the *vocabulary* of `Word` to any *collections* and *dictionaries*.

### 4. Answer

```go
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

const (
	Collection = 1
	Dictionary = 2
	Online = 3
	Basic = "basic"
	Advance = "advance"
)
```

This is the data structure of the *vocabulary*.

#### 4.1. SourceName

`SourceName` indicates where the *vocabulary* comes from. It can be the name of:

- the *collection*


- the *dictionary*
- the *service*

#### 4.2. Type

- If the vocabulary comes from the *collection*, its `type` will be `Collection`.
- If the vocabulary comes from the *dictionary*, its `type` will be `Dictionary`.
- If the vocabulary comes from the *service*, its `type` will be `Online`.

#### 4.3. Status

`Status` indicates some other information.

If the *vocabulary* comes from the *collection* or the *dictionary*:

- If the *vocabulary* is in the *collection* and the *dictionary*, its `Status` will be `""`.
- If the *vocabulary* comes from *basic query*, its `Status` will be `Basic`.
- If the *vocabulary* comes from *advanced query*, its `Status` will be `Advance`.

If the *vocabulary* comes from the *service*, its `Status` will be `Basic`.

#### 4.4. Location

`Location` indicates the *vocabulary's* location. It is a structure and has got these members:

- Integer `TableIndex`: it indicates the list of *collection* or *dictionary* (depends on `Type`) which the *vocabulary* belongs to.
- Integer `ItemIndex`: it indicates the *vocabulary's* index in the *collection* or the *dictionary*.

**These indexes begin with 0.**

`Location` can be used to modify contents of *collections* and *dictionaries*.

### 5. Edit

```go
type VocabularyEditStruct struct {
	TableType int `json:"tableType"`
	Location  struct {
		TableIndex int `json:"tableIndex"`
		ItemIndex  int `json:"itemIndex"`
	} `json:"location"`
	ContentType int    `json:"contentType"`
	Content     string `json:"content"`
}

const (
	Define = -1
	Note = -2
)
```

This structure is used to edit a *vocabulary* in *collection* or *dictionary*.

#### 5.1. TableType

`TableType` is same as `Type` in #4.2.

#### 5.2. Location

`Location` is same as #4.4.

#### 5.3. ContentType

- To edit `Define` of the *vocabulary*, `ContentType` should be set as `Define`.
- To edit `Note` of the *vocabulary*, `ContentType` should be set as `Note`.

#### 5.4. Content

`Content` is the value of editing.

### 6. Others

- All ".md" files are edited by [Typora](http://typora.io).
- The style of all ".md" files is [Github Flavored Markdown](https://guides.github.com/features/mastering-markdown/#GitHub-flavored-markdown).
- There is a LF (Linux) at the end of each line.
