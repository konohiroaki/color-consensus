package consensus

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ColorConsensus struct {
	// https://ja.wikipedia.org/wiki/ISO_639-1
	// https://godoc.org/golang.org/x/text/language
	Lang      string `json:"lang" bson:"lang"`
	Name      string `json:"name" bson:"name"`
	Code      string `json:"code" bson:"code"`
	VoteCount int    `json:"vote" bson:"vote"`
	//TODO: validate keys is v9 feature but gin still uses v8 validator.
	// https://github.com/gin-gonic/gin/pull/1015
	Colors map[string]int `json:"colors" bson:"colors" validate:"dive,keys,hexcolor,endkeys"`
}

var consensusCollection *mgo.Collection

func InitRepo(uri, db string) {
	session, _ := mgo.Dial(uri)
	c := session.DB(db).C("consensus")
	consensusCollection = c
}

func Get(lang, name string) (ColorConsensus, bool) {
	var consensus ColorConsensus
	if err := consensusCollection.Find(bson.M{"lang": lang, "name": name}).One(&consensus); err != nil {
		return ColorConsensus{}, false
	}
	return consensus, true
}

func GetList() []ColorConsensus {
	var userList []ColorConsensus
	_ = consensusCollection.Find(bson.M{}).All(&userList)
	return userList
}

func GetKeys() []interface{} {
	var keyList []interface{}
	_ = consensusCollection.Find(nil).Select(bson.M{"lang": 1, "name": 1, "code": 1}).All(&keyList)
	if keyList == nil {
		return []interface{}{}
	}
	return keyList
}

func Add(consensus ColorConsensus) {
	consensus.Colors = map[string]int{}
	consensus.VoteCount = 0

	_ = consensusCollection.Insert(&consensus)
}

func Update(lang, name string, colorsToIncrement, colorsToDecrement []string) {
	incrementMap := map[string]int{}
	if len(colorsToIncrement) > 0 && len(colorsToDecrement) == 0 {
		incrementMap["vote"] = 1
	} else if len(colorsToIncrement) == 0 && len(colorsToDecrement) > 0 {
		incrementMap["vote"] = -1
	}

	for i := 0; i < len(colorsToIncrement); i++ {
		incrementMap["colors."+colorsToIncrement[i]] = 1
	}
	for i := 0; i < len(colorsToDecrement); i++ {
		key := "colors." + colorsToDecrement[i]
		if _, exist := incrementMap[key]; exist {
			incrementMap[key] = 0
		} else {
			incrementMap[key] = -1
		}
	}

	_, _ = consensusCollection.Upsert(bson.M{"lang": lang, "name": name}, bson.M{"$inc": incrementMap})
}

func InsertSampleData() {
	votes := []*ColorConsensus{
		{Lang: "en", Name: "red", Code: "#ff0000", VoteCount: 2, Colors: map[string]int{"#ff0000": 1, "#f00000": 1}},
		{Lang: "en", Name: "green", Code: "#008000", VoteCount: 1, Colors: map[string]int{"#008000": 1}},
		{Lang: "ja", Name: "赤", Code: "#bf1e33", VoteCount: 1, Colors: map[string]int{"#bf1e33": 1}},
		{Lang: "en", Name: "gray", Code: "#808080", VoteCount: 0, Colors: map[string]int{}},
	}

	_, _ = consensusCollection.RemoveAll(nil)
	tmp := []interface{}{}
	for _, v := range votes {
		tmp = append(tmp, v)
	}
	_ = consensusCollection.Insert(tmp...)
}
