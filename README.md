# Vocabulary for MYDICTIONARY

[简体中文版](./README.zh-Hans.md)

### 1. Introduction

The repository defines some basic data structures of [MYDICTIONARY](https://github.com/zzc-tongji/mydictionary) ("the core library" for short).

### 2. Query

Take file `animal.xlsx` as an example.

![animal](./README.picture/animal.png)

#### 2.1. Basic Query (Exact Match)

If the word **is exactly the same as the content in cell at column "Word" in a line**, data of the line will be selected.

For example, in *basic query*, input word "cat" and data of the line 1 will be selected.

#### 2.2. Advanced Query (Keyword Match)

If the word **is contained by the content in cells at columns "Word", "Define" or "Note" in a line**, the line will be selected.

For example, in *advanced query*, input word "e" and the data of line 1, 3 and 5 will be selected.

### 3. VocabularyAskStruct

```go
type VocabularyAskStruct struct {
	Word        string `json:"word"`
	Advance     bool   `json:"advance"`
	Online      bool   `json:"online"`
	DoNotRecord bool   `json:"doNotRecord"`
}
```

This structure indicates the word and options of query.

#### 3.1. Word

`Word` indicates the word of query.

#### 3.2. Advance

If `Advance` is `false`, the core library will do *basic query* of `Word` only.

If `Advance` is `true`, the core library will do both *basic query* and *advanced query* of `Word`.

#### 3.3. Online

If `Online` is `true`, the core library will know that users need query `Word` online.

**Note that it doesn't ensure that the core library will query `Word` online.** Whether the core library does also depends on `online.mode` in *configuration* (at [here](https://github.com/zzc-tongji/mydictionary#2431-mode)).

#### 3.4. DoNotRecord

If `DoNotRecord` is `true`, the core library will not *record* the *vocabulary* to any *collections* and *dictionaries*.

### 4. VocabularyAnswerStruct

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

#### 4.1. Word

`Word` indicates the word in the *vocabulary*.

#### 4.2. Define

`Define` indicates definitions in the *vocabulary*.

#### 4.3. SerialNumber

`SerialNumber` indicates the serial number of the *vocabulary*.

#### 4.4. QueryCounter

`QueryCounter` indicates the query counter of the *vocabulary*.

#### 4.5. Note

`Note` indicates notes in the *vocabulary*.

#### 4.6. QueryTime

`QueryTime` indicates the last query time of the *vocabulary*.

#### 4.7. SourceName

`SourceName` indicates where the *vocabulary* comes from. It can be the name of:

- the *collection*
- the *dictionary*
- the *service*

#### 4.8. Type

- If the *vocabulary* comes from the *collection*, its `Type` will be `Collection`.
- If the *vocabulary* comes from the *dictionary*, its `Type` will be `Dictionary`.
- If the *vocabulary* comes from the *service*, its `Type` will be `Online`.

#### 4.9. Status

`Status` indicates some other information.

If the *vocabulary* comes from the *collection* or the *dictionary*:

- If the *vocabulary* comes from *basic query*, its `Status` will be `Basic`.
- If the *vocabulary* comes from *advanced query*, its `Status` will be `Advance`.

If the *vocabulary* comes from the *service*, its `Status` will be `Basic`.

#### 4.10. Location

`Location` is used to locate the *vocabulary* in *collections* or *dictionaries*.

It is a structure and has got these members:

- `TableIndex` indicates the index of the *collection* or the *dictionary* (depends on `Type`) which the *vocabulary* belongs to in the list.
- `ItemIndex` indicates the *vocabulary's* index in the *collection* or the *dictionary*.

**These indexes begin with 0.**

### 5. VocabularyEditStruct

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

This structure is used to provide information for editting a *vocabulary* in *collection* or *dictionary*.

#### 5.1. TableType

`TableType` is same as `Type` in \#4.8.

#### 5.2. Location

`Location` is same as \#4.10.

#### 5.3. ContentType

- To edit `Define` of the *vocabulary*, `ContentType` should be set as `Define`.
- To edit `Note` of the *vocabulary*, `ContentType` should be set as `Note`.

#### 5.4. Content

`Content` is the new string to be set.

### 6. Others

- All code files are edited by [Atom](https://atom.io/).
- All ".md" files are edited by [Typora](http://typora.io).
- The style of all ".md" files is [Github Flavored Markdown](https://guides.github.com/features/mastering-markdown/#GitHub-flavored-markdown).
- There is a LF (Linux) at the end of each line.
