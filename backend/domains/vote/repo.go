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
	return voteList
}

func FindList(lang, name string) []ColorVote {
	var voteList []ColorVote
	_ = voteCollection.Find(bson.M{"lang": lang, "name": name}).All(&voteList)
	return voteList
}

// TODO: check consensus document existence before update
// TODO: do transaction management with mgo/txn
// TODO: if there's already same lang,name,user document, overwrite it.
func Add(vote ColorVote) bool {
	vote.Date = time.Now()
	_ = voteCollection.Insert(&vote)
	consensus.Update(vote.Language, vote.ColorName, vote.Colors)
	return true
}

func InsertSampleData() {
	votes := []*ColorVote{
		{Language: "en", ColorName: "red", User: "foo", Date: time.Now(), Colors: []string{"#ff0000", "#ff007f"}},
		{Language: "en", ColorName: "red", User: "bar", Date: time.Now(), Colors: []string{"#ff0000", "#ff007f"}},
		{Language: "en", ColorName: "red", User: "baz", Date: time.Now(), Colors: []string{"#ff0000", "#ff007f"}},
		{Language: "en", ColorName: "red", User: "aaa", Date: time.Now(), Colors: []string{"#ff0000", "#ff007f"}},
	}

	_, _ = voteCollection.RemoveAll(nil)
	tmp := []interface{}{}
	for _, v := range votes {
		tmp = append(tmp, v)
	}
	_ = voteCollection.Insert(tmp...)
}
