package repositories

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

type ColorRepository interface {
	Add(lang, name, code, user string)
	GetAll(fields []string) []map[string]interface{}
}

type colorRepository struct {
	Collection *mgo.Collection
}

func NewColorRepository(env string) ColorRepository {
	uri, db := getDatabaseURIAndName()
	session, _ := mgo.Dial(uri)
	collection := session.DB(db).C("color")
	repository := &colorRepository{collection}

	if env == "development" {
		repository.insertSampleData()
	}

	return repository
}

type color struct {
	Lang string    `bson:"lang"`
	Name string    `bson:"name"`
	Code string    `bson:"code"`
	User string    `bson:"user"`
	Date time.Time `bson:"date"`
}

func (r colorRepository) Add(lang, name, code, user string) {
	err := r.Collection.Insert(color{
		Lang: lang,
		Name: name,
		Code: code,
		User: user,
	})
	if err != nil {
		log.Println(err)
	}
}

func (r colorRepository) GetAll(fields []string) []map[string]interface{} {
	var result []map[string]interface{}
	err := r.Collection.
		Pipe([]bson.M{{"$project": r.getProjector(fields)}}).
		All(&result)

	if result == nil {
		if err != nil {
			log.Println(err)
		}
		return []map[string]interface{}{}
	}
	return result
}

func (r colorRepository) getProjector(fields []string) bson.M {
	var projector = bson.M{}
	for _, field := range fields {
		projector[field] = 1;
	}
	projector["_id"] = 0;
	return projector
}

func (r colorRepository) insertSampleData() {
	votes := []*color{
		{Lang: "en", Name: "red", User: "00943efe-0aa5-46a4-ae5b-6ef818fc1480", Code: "#ff0000", Date: time.Now()},
		{Lang: "en", Name: "lime", User: "0da04f70-dc71-4674-b47b-365c3b0805c4", Code: "#00ff00", Date: time.Now()},
		{Lang: "en", Name: "green", User: "0da04f70-dc71-4674-b47b-365c3b0805c4", Code: "#008000", Date: time.Now()},
		{Lang: "ja", Name: "赤", User: "20af3406-8c7e-411a-851f-31732416fa83", Code: "#bf1e33", Date: time.Now()},
		{Lang: "en", Name: "gray", Code: "#808080", Date: time.Now()},
	}

	_, _ = r.Collection.RemoveAll(nil)
	tmp := []interface{}{}
	for _, v := range votes {
		tmp = append(tmp, v)
	}
	_ = r.Collection.Insert(tmp...)
}
