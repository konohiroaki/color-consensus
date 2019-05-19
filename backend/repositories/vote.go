package repositories

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"log"
	"sync"
	"time"
)

type VoteRepository interface {
	Add(lang, name string, newColors []string, userID string)
	Get(lang, name string, fields []string) []map[string]interface{}
	RemoveByUser(userID string)
}

type voteRepository struct {
	Collection *mgo.Collection
}

var (
	voteRepoInstance VoteRepository
	voteRepoOnce     sync.Once
)

func GetVoteRepository(env string) VoteRepository {
	voteRepoOnce.Do(func() {
		uri, name := getDatabaseURIAndName()
		session, _ := mgo.Dial(uri)
		database := session.DB(name)
		repository := newVoteRepository(database)

		if env == "development" {
			repository.insertSampleData()
		}
		voteRepoInstance = repository
	})
	return voteRepoInstance
}

func newVoteRepository(database *mgo.Database) *voteRepository {
	return &voteRepository{database.C("vote")}
}

type vote struct {
	Lang   string    `bson:"lang"`
	Name   string    `bson:"name"`
	Colors []string  `bson:"colors"`
	Date   time.Time `bson:"date"`
	User   string    `bson:"user"`
}

func (r voteRepository) Add(lang, name string, newColors []string, userID string) {
	vote := vote{Lang: lang, Name: name, Colors: newColors, Date: time.Now(), User: userID}
	_, _ = r.Collection.Upsert(bson.M{"lang": lang, "name": name, "user": userID}, &vote)
}

func (r voteRepository) Get(lang, name string, fields []string) []map[string]interface{} {
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

func (r voteRepository) RemoveByUser(userID string) {
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
	// ageGroup could be wrong since user input is only for year, but it's small problem. :D
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
	votes := []*vote{
		{Lang: "en", Name: "red", Colors: []string{"#ff0000"}, Date: time.Now(), User: "00943efe-0aa5-46a4-ae5b-6ef818fc1480"},
		{Lang: "en", Name: "red", Colors: []string{"#f00000"}, Date: time.Now(), User: "0da04f70-dc71-4674-b47b-365c3b0805c4"},
		{Lang: "en", Name: "green", Colors: []string{"#008000"}, Date: time.Now(), User: "0da04f70-dc71-4674-b47b-365c3b0805c4"},
		{Lang: "ja", Name: "赤", Colors: []string{"#bf1e33"}, Date: time.Now(), User: "20af3406-8c7e-411a-851f-31732416fa83"},
	}

	_, _ = r.Collection.RemoveAll(nil)
	tmp := []interface{}{}
	for _, v := range votes {
		tmp = append(tmp, v)
	}
	_ = r.Collection.Insert(tmp...)
}
