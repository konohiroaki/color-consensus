package models

type ColorConsensus struct {
	// https://ja.wikipedia.org/wiki/ISO_639-1
	// https://godoc.org/golang.org/x/text/language
	Language string `json:"lang"`
	Color    string `json:"name"`
	Code     string `json:"code"`
	Vote     int    `json:"vote"`
	//TODO: validate keys is v9 feature but gin still uses v8 validator.
	// https://github.com/gin-gonic/gin/pull/1015
	Colors map[string]int `json:"colors" validate:"dive,keys,hexcolor,endkeys"`
}
