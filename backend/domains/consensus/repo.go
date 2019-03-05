package consensus

import (
	"github.com/konohiroaki/color-consensus/backend/config"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ColorConsensus struct {
	// https://ja.wikipedia.org/wiki/ISO_639-1
	// https://godoc.org/golang.org/x/text/language
	Language  string `json:"lang" bson:"lang"`
	ColorName string `json:"name" bson:"name"`
	ColorCode string `json:"code" bson:"code"`
	VoteCount int    `json:"vote" bson:"vote"`
	//TODO: validate keys is v9 feature but gin still uses v8 validator.
	// https://github.com/gin-gonic/gin/pull/1015
	Colors map[string]int `json:"colors" bson:"colors" validate:"dive,keys,hexcolor,endkeys"`
}

var consensusCollection *mgo.Collection

func InitRepo() {
	uri := config.GetConfig().Get("mongo.url").(string)
	session, _ := mgo.Dial(uri)
	c := session.DB("cc").C("consensus")
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
	return keyList
}

func Add(consensus ColorConsensus) {
	consensus.Colors = map[string]int{}
	consensus.VoteCount = 0

	_ = consensusCollection.Insert(&consensus)
}

// TODO: should receive removed vote as well when vote is overwriting another vote.
func Update(lang, name string, add []string) {
	colorMap := map[string]int{}
	for i := 0; i < len(add); i++ {
		colorMap["add."+add[i]] = 1
	}
	_, _ = consensusCollection.Upsert(bson.M{"lang": lang, "name": name}, bson.M{"$inc": colorMap})
}

func InsertSampleData() {
	votes := []*ColorConsensus{
		{Language: "en", ColorName: "red", ColorCode: "#ff0000", VoteCount: 10, Colors: map[string]int{"#ff0000": 10, "#ff007f": 3}},
		{Language: "en", ColorName: "green", ColorCode: "#008000", VoteCount: 10, Colors: map[string]int{"#00ff00": 10, "#00ff33": 3}},
		{Language: "ja", ColorName: "赤", ColorCode: "#bf1e33", VoteCount: 15, Colors: map[string]int{"#ff0000": 15, "#ff007f": 5}},
		{Language: "en", ColorName: "gray", ColorCode: "#808080", VoteCount: 0, Colors: map[string]int{}},
	}

	_, _ = consensusCollection.RemoveAll(nil)
	tmp := []interface{}{}
	for _, v := range votes {
		tmp = append(tmp, v)
	}
	_ = consensusCollection.Insert(tmp...)
}
