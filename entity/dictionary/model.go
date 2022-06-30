package dictionary

type DocOutPut []*Docs

type Docs struct {
	Word      string      `json:"word"`
	Phonetics []Phonetics `json:"phonetics"`
	Meanings  []Meanings  `json:"meanings"`
}

type Phonetics struct {
	Text  string `json:"text"`
	Audio string `json:"audio,omitempty"`
}

type Meanings struct {
	PartOfSpeech string        `json:"partOfSpeech"`
	Definitions  []Definitions `json:"definitions"`
}

type Definitions struct {
	Definition string `json:"definition"`
	Example    string `json:"example"`
}
