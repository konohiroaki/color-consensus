package dto

import "time"

type ColorConsensus struct {
	// https://ja.wikipedia.org/wiki/ISO_639-1
	// https://godoc.org/golang.org/x/text/language
	Language string `json:"-"`
	Color    string `json:"-"`
	Vote     int    `json:"vote"`
	//TODO: validate keys is v9 feature but gin still uses v8 validator.
	// https://github.com/gin-gonic/gin/pull/1015
	Colors map[string]int `json:"colors" validate:"dive,keys,hexcolor,endkeys"`
}

type ColorVote struct {
	Language string    `json:"-"`
	Color    string    `json:"-"`
	User     string    `json:"-"`
	Date     time.Time `json:"date"`
	//TODO: not working when POST.
	Colors []string `json:"colors" validate:"dive,hexcolor"`
}

/*
 * In DB, associate this data with cookie data. Also, generate a key to reuse that user data when cookie is gone.
 * If user have data, just go to the input screen.
 * User should be able to list input history in some way.
 */
type User struct {
	ID string `json:"id"`
	// https://ja.wikipedia.org/wiki/ISO_3166-1
	Nationality string    `json:"nationality"`
	Gender      string    `json:"gender"`
	Birth       uint      `json:"birth"`
	Date        time.Time `json:"date"`
}

var Sum = []*ColorConsensus{
	{Language: "en", Color: "red", Vote: 10, Colors: map[string]int{"ff0000": 10, "#ff007f": 3}},
	{Language: "en", Color: "green", Vote: 10, Colors: map[string]int{"00ff00": 10, "#00ff33": 3}},
	{Language: "ja", Color: "赤", Vote: 15, Colors: map[string]int{"ff0000": 15, "#ff007f": 5}},
}

var Raw = []*ColorVote{
	{Language: "en", Color: "red", User: "foo", Date: time.Now(), Colors: []string{"ff0000", "#ff007f"}},
	{Language: "en", Color: "red", User: "bar", Date: time.Now(), Colors: []string{"#ff0000", "#ff007f"}},
	{Language: "en", Color: "red", User: "baz", Date: time.Now(), Colors: []string{"#ff0000", "#ff007f"}},
	{Language: "en", Color: "red", User: "aaa", Date: time.Now(), Colors: []string{"#ff0000", "#ff007f"}},
}

var Users = []*User{
	{ID: "0da04f70-dc71-4674-b47b-365c3b0805c4", Nationality: "Japan", Gender: "Male", Birth: 1990, Date: time.Now()},
	{ID: "20af3406-8c7e-411a-851f-31732416fa83", Nationality: "Japan", Gender: "Male", Birth: 1991, Date: time.Now()},
}
