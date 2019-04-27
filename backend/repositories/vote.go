package repositories

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

type VoteRepository interface {
	Add(user, lang, name string, newColors []string)
	GetVotes(lang, name string, fields []string) []map[string]interface{}
	RemoveForUser(userID string)
}

type voteRepository struct {
	Collection *mgo.Collection
}

func NewVoteRepository(uri, db, env string) VoteRepository {
	session, _ := mgo.Dial(uri)
	collection := session.DB(db).C("vote")
	repository := &voteRepository{collection}

	if env == "development" {
		log.Println("detected development mode. inserting sample vote data.")
		repository.insertSampleData()
	}

	return repository
}

type colorVote struct {
	Lang   string    `bson:"lang"`
	Name   string    `bson:"name"`
	User   string    `bson:"user"`
	Date   time.Time `bson:"date"`
	Colors []string  `bson:"colors"`
}

func (r voteRepository) Add(user, lang, name string, newColors []string) {
	vote := colorVote{Lang: lang, Name: name, User: user, Date: time.Now(), Colors: newColors}
	_, _ = r.Collection.Upsert(bson.M{"lang": lang, "name": name, "user": user}, &vote)
}

func (r voteRepository) GetVotes(lang, name string, fields []string) []map[string]interface{} {
	var result []map[string]interface{}
	err := r.Collection.
		Pipe(r.getAggregators(lang, name, fields)).
		All(&result)

	if result == nil {
		if err != nil {
			log.Println(err)
		}
		return []map[string]interface{}{}
	}
	return result
}

func (r voteRepository) RemoveForUser(userID string) {
	_, _ = r.Collection.RemoveAll(bson.M{"user": userID})
}

func (r voteRepository) getAggregators(lang, name string, fields []string) []bson.M {
	var aggregators = []bson.M{}
	aggregators = append(aggregators, bson.M{"$match": r.getMatcher(lang, name)})
	aggregators = append(aggregators, r.getUserLookUpAggregators()...)
	aggregators = append(aggregators, bson.M{"$project": r.getProjector(fields)})
	return aggregators
}

func (r voteRepository) getUserLookUpAggregators() []bson.M {
	// ageGroup maybe wrong since user input is only for year.
	// ageGroup = Math.floor((currentYear - birthYear) / 10) * 10
	ageGroupAggregator :=
		bson.M{"$multiply": []interface{}{
			bson.M{"$floor":
			bson.M{"$divide": []interface{}{
				bson.M{"$subtract": []interface{}{bson.M{"$year": time.Now()}, "$voter.birth"}},
				10}}}, 10}};
	return []bson.M{
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
			"voter.ageGroup":    ageGroupAggregator,
			"voter.gender":      1,
			"lang":              1,
			"name":              1,
			"colors":            1,
			"date":              1,
		}},
	}
}

func (r voteRepository) getMatcher(lang, name string) bson.M {
	var matcher = bson.M{}
	r.setIfValuePresent(matcher, "lang", lang)
	r.setIfValuePresent(matcher, "name", name)
	return matcher
}

func (r voteRepository) setIfValuePresent(m map[string]interface{}, key, value string) {
	if (value != "") {
		m[key] = value
	}
}

func (r voteRepository) getProjector(fields []string) bson.M {
	var projector = bson.M{}
	for _, field := range fields {
		projector[field] = 1;
	}
	return projector
}

func (r voteRepository) insertSampleData() {
	votes := []*colorVote{
		{Lang: "en", Name: "red", User: "00943efe-0aa5-46a4-ae5b-6ef818fc1480", Date: time.Now(), Colors: []string{"#ff0000"}},
		{Lang: "en", Name: "red", User: "0da04f70-dc71-4674-b47b-365c3b0805c4", Date: time.Now(), Colors: []string{"#f00000"}},
		{Lang: "en", Name: "green", User: "0da04f70-dc71-4674-b47b-365c3b0805c4", Date: time.Now(), Colors: []string{"#008000"}},
		{Lang: "ja", Name: "赤", User: "20af3406-8c7e-411a-851f-31732416fa83", Date: time.Now(), Colors: []string{"#bf1e33"}},
	}

	_, _ = r.Collection.RemoveAll(nil)
	tmp := []interface{}{}
	for _, v := range votes {
		tmp = append(tmp, v)
	}
	_ = r.Collection.Insert(tmp...)
}
