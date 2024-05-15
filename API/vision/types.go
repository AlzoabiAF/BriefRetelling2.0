package vision

type body struct {
	MineTime      string   `json:"mineTime"`
	LanguageCodes []string `json:"languageCodes"`
	Model         string   `json:"model"`
	Content       string   `json:"content"`
}

type vertices struct {
	X string `json:"x"`
	Y string `json:"y"`
}

type boundingBox struct {
	Vertices []vertices `json:"vertices"`
}

type textSegments struct {
	StartIndex string `json:"startIndex"`
	Length     string `json:"length"`
}

type words struct {
	BoundingBox  boundingBox    `json:"boundingBox"`
	Text         string         `json:"text"`
	EntityIndex  string         `json:"entityIndex"`
	TextSegments []textSegments `json:"textSegments"`
}

type lines struct {
	BoundingBox  boundingBox    `json:"boundingBox"`
	Text         string         `json:"text"`
	Words        []words        `json:"words"`
	TextSegments []textSegments `json:"textSegments"`
	Orientation  string         `json:"orientation"`
}

type languages struct {
	LanguageCode string `json:"languageCode"`
}

type blocks struct {
	BoundingBox  boundingBox    `json:"boundingBox"`
	Lines        []lines        `json:"lines"`
	Languages    []languages    `json:"languages"`
	TextSegments []textSegments `json:"textSegments"`
}

type textAnnotation struct {
	Width    string        `json:"width"`
	Height   string        `json:"height"`
	Blocks   []blocks      `json:"blocks"`
	Entities []interface{} `json:"entities"`
	Tables   []interface{} `json:"tables"`
	FullText string        `json:"fullText"`
	Rotate   string        `json:"rotate"`
}

type result struct {
	TextAnnotation textAnnotation `json:"textAnnotation"`
	Page           string         `json:"page"`
}

type response struct {
	Result result `json:"result"`
}
