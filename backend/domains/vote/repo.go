package vote

import (
	"github.com/konohiroaki/color-consensus/backend/domains/consensus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type ColorVote struct {
	Language  string    `json:"lang" bson:"lang"`
	ColorName string    `json:"name" bson:"name"`
	User      string    `json:"user" bson:"user"`
	Date      time.Time `json:"date" bson:"date"`
	//FIXME: validate not working.
	Colors []string `json:"colors" bson:"colors" validate:"dive,hexcolor"`
}

var voteCollection *mgo.Collection

func InitRepo(uri, db string) {
	session, _ := mgo.Dial(uri)
	c := session.DB(db).C("vote")
	voteCollection = c
}

func GetList() []ColorVote {
	var voteList []ColorVote
	_ = voteCollection.Find(bson.M{}).All(&voteList)
	if voteList == nil {
		return []ColorVote{}
	}
	return voteList
}

func FindList(lang, name string) []ColorVote {
	var voteList []ColorVote
	_ = voteCollection.Find(bson.M{"lang": lang, "name": name}).All(&voteList)
	if voteList == nil {
		return []ColorVote{}
	}
	return voteList
}

// TODO: check consensus document existence before update
// TODO: do transaction management with mgo/txn
// TODO: if there's already same lang,name,user document, overwrite it.
func Add(vote ColorVote) bool {
	vote.Date = time.Now()
	_ = voteCollection.Insert(&vote)
	consensus.Update(vote.Language, vote.ColorName, vote.Colors, []string{})
	return true
}

func RemoveForUser(userID string) {
	var votes []ColorVote
	_ = voteCollection.Find(bson.M{"user": userID}).All(&votes)
	for _, vote := range votes {
		consensus.Update(vote.Language, vote.ColorName, []string{}, vote.Colors)
	}
	_, _ = voteCollection.RemoveAll(bson.M{"user": userID})
}

func InsertSampleData() {
	votes := []*ColorVote{
		{Language: "en", ColorName: "red", User: "00943efe-0aa5-46a4-ae5b-6ef818fc1480", Date: time.Now(), Colors: []string{"#ff0000"}},
		{Language: "en", ColorName: "green", User: "0da04f70-dc71-4674-b47b-365c3b0805c4", Date: time.Now(), Colors: []string{"#008000"}},
		{Language: "ja", ColorName: "赤", User: "20af3406-8c7e-411a-851f-31732416fa83", Date: time.Now(), Colors: []string{"#bf1e33"}},
		{Language: "en", ColorName: "red", User: "20af3406-8c7e-411a-851f-31732416fa83", Date: time.Now(), Colors: []string{"#f00000"}},
	}

	_, _ = voteCollection.RemoveAll(nil)
	tmp := []interface{}{}
	for _, v := range votes {
		tmp = append(tmp, v)
	}
	_ = voteCollection.Insert(tmp...)
}
