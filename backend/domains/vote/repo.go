package vote

import (
	"fmt"
	"github.com/konohiroaki/color-consensus/backend/domains/consensus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type colorVote struct {
	Lang string    `bson:"lang"`
	Name string    `bson:"name"`
	User string    `bson:"user"`
	Date time.Time `bson:"date"`
	//FIXME: validate not working.
	Colors []string `bson:"colors" validate:"dive,hexcolor"`
}

var voteCollection *mgo.Collection

var userLookup = []bson.M{
	{"$lookup": bson.M{
		"from":         "user",
		"localField":   "user",
		"foreignField": "id",
		"as":           "voter",
	}},
	{"$unwind": "$voter"},
	{"$project": bson.M{
		"_id":               0,
		"voter.nationality": 1,
		"voter.gender":      1,
		"voter.birth":       1,
		"lang":              1,
		"name":              1,
		"colors":            1,
		"date":              1,
	}},
}

func InitRepo(uri, db string) {
	session, _ := mgo.Dial(uri)
	c := session.DB(db).C("vote")
	voteCollection = c
}

func GetVotes(lang, name string, fields []string) []bson.M {
	var result []bson.M
	err := voteCollection.
		Pipe(getAggregators(lang, name, fields)).
		All(&result)

	if result == nil {
		if err != nil {
			fmt.Println(err)
		}
		return []bson.M{}
	}
	return result
}

func getAggregators(lang, name string, fields []string) []bson.M {
	var aggregators = []bson.M{}
	aggregators = append(aggregators, bson.M{"$match": getMatcher(lang, name)})
	aggregators = append(aggregators, userLookup...)
	aggregators = append(aggregators, bson.M{"$project": getProjector(fields)})
	return aggregators
}

func getMatcher(lang, name string) bson.M {
	var matcher = bson.M{}
	setIfValuePresent(matcher, "lang", lang)
	setIfValuePresent(matcher, "name", name)
	return matcher
}

func setIfValuePresent(m map[string]interface{}, key, value string) {
	if (value != "") {
		m[key] = value
	}
}

func getProjector(fields []string) bson.M {
	var projector = bson.M{}
	for _, field := range fields {
		projector[field] = 1;
	}
	return projector
}

// TODO: do transaction management with mgo/txn?
func Add(user, lang, name string, newColors []string) bool {
	oldColors := getOldVoteColors(user, lang, name)

	vote := colorVote{Lang: lang, Name: name, User: user, Date: time.Now(), Colors: newColors}
	_, _ = voteCollection.Upsert(bson.M{"lang": lang, "name": name, "user": user}, &vote)

	// maybe we can remove consensus collection. it's just an aggregated result of votes.
	consensus.Update(lang, name, newColors, oldColors)
	return true
}

func getOldVoteColors(user, lang, name string) []string {
	var old colorVote
	err := voteCollection.Find(bson.M{"lang": lang, "name": name, "user": user}).
		Select(bson.M{"colors": 1}).One(&old)

	if err == nil {
		return old.Colors
	} else {
		return []string{}
	}
}

func RemoveForUser(userID string) {
	var votes []colorVote
	_ = voteCollection.Find(bson.M{"user": userID}).All(&votes)
	for _, vote := range votes {
		consensus.Update(vote.Lang, vote.Name, []string{}, vote.Colors)
	}
	_, _ = voteCollection.RemoveAll(bson.M{"user": userID})
}

func InsertSampleData() {
	votes := []*colorVote{
		{Lang: "en", Name: "red", User: "00943efe-0aa5-46a4-ae5b-6ef818fc1480", Date: time.Now(), Colors: []string{"#ff0000"}},
		{Lang: "en", Name: "green", User: "0da04f70-dc71-4674-b47b-365c3b0805c4", Date: time.Now(), Colors: []string{"#008000"}},
		{Lang: "ja", Name: "赤", User: "20af3406-8c7e-411a-851f-31732416fa83", Date: time.Now(), Colors: []string{"#bf1e33"}},
		{Lang: "en", Name: "red", User: "20af3406-8c7e-411a-851f-31732416fa83", Date: time.Now(), Colors: []string{"#f00000"}},
	}

	_, _ = voteCollection.RemoveAll(nil)
	tmp := []interface{}{}
	for _, v := range votes {
		tmp = append(tmp, v)
	}
	_ = voteCollection.Insert(tmp...)
}
