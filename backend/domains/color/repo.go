package color

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type color struct {
	Lang string    `bson:"lang"`
	Name string    `bson:"name"`
	Code string    `bson:"code"`
	User string    `bson:"user"`
	Date time.Time `bson:"date"`
}

var colorCollection *mgo.Collection

func InitRepo(uri, db string) {
	session, _ := mgo.Dial(uri)
	c := session.DB(db).C("color")
	colorCollection = c
}

func GetAll(fields []string) []bson.M {
	var result []bson.M
	err := colorCollection.
		Pipe([]bson.M{{"$project": getProjector(fields)}}).
		All(&result)

	if result == nil {
		if err != nil {
			fmt.Println(err)
		}
		return []bson.M{}
	}
	return result
}

func getProjector(fields []string) bson.M {
	var projector = bson.M{}
	for _, field := range fields {
		projector[field] = 1;
	}
	projector["_id"] = 0;
	return projector
}

func InsertSampleData() {
	votes := []*color{
		{Lang: "en", Name: "red", User: "00943efe-0aa5-46a4-ae5b-6ef818fc1480", Code: "#ff0000", Date: time.Now()},
		{Lang: "en", Name: "lime", User: "0da04f70-dc71-4674-b47b-365c3b0805c4", Code: "#00ff00", Date: time.Now()},
		{Lang: "en", Name: "green", User: "0da04f70-dc71-4674-b47b-365c3b0805c4", Code: "#008000", Date: time.Now()},
		{Lang: "ja", Name: "赤", User: "20af3406-8c7e-411a-851f-31732416fa83", Code: "#bf1e33", Date: time.Now()},
		{Lang: "en", Name: "gray", Code: "#808080", Date: time.Now()},
	}

	_, _ = colorCollection.RemoveAll(nil)
	tmp := []interface{}{}
	for _, v := range votes {
		tmp = append(tmp, v)
	}
	_ = colorCollection.Insert(tmp...)
}
