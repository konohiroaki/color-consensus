package repositories

import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"log"
	"sync"
	"time"
)

type ColorRepository interface {
	Add(category, name, code, user string) error
	GetAll(fields []string) []map[string]interface{}
}

type colorRepository struct {
	Collection *mgo.Collection
}

var (
	colorRepoInstance ColorRepository
	colorRepoOnce     sync.Once
)

func GetColorRepository(env string) ColorRepository {
	colorRepoOnce.Do(func() {
		uri, name := getDatabaseURIAndName()
		session, _ := mgo.Dial(uri)
		database := session.DB(name)
		repository := newColorRepository(database)

		if env == "development" {
			repository.insertSampleData()
		}
		colorRepoInstance = repository
	})
	return colorRepoInstance
}

func newColorRepository(database *mgo.Database) *colorRepository {
	return &colorRepository{database.C("color")}
}

type color struct {
	Category string    `bson:"category"`
	Name     string    `bson:"name"`
	Code     string    `bson:"code"`
	User     string    `bson:"user"`
	Date     time.Time `bson:"date"`
}

func (r colorRepository) Add(category, name, code, userID string) error {
	count, _ := r.Collection.Find(bson.M{"category": category, "name": name}).Limit(1).Count()
	if count != 0 {
		return fmt.Errorf("the requested color already exists")
	}
	err := r.Collection.Insert(color{
		Category: category,
		Name:     name,
		Code:     code,
		User:     userID,
		Date:     time.Now(),
	})
	if err != nil {
		log.Println(err)
	}
	return err
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
		{Category: "HTML Color", Name: "Red", Code: "#ff0000", User: "00943efe-0aa5-46a4-ae5b-6ef818fc1480", Date: time.Now()},
		{Category: "X11", Name: "Green", Code: "#00ff00", User: "0da04f70-dc71-4674-b47b-365c3b0805c4", Date: time.Now()},
		{Category: "HTML Color", Name: "Green", Code: "#008000", User: "0da04f70-dc71-4674-b47b-365c3b0805c4", Date: time.Now()},
		{Category: "JIS慣用色名", Name: "赤", Code: "#bf1e33", User: "20af3406-8c7e-411a-851f-31732416fa83", Date: time.Now()},
		{Category: "HTML Color", Name: "Gray", Code: "#808080", User: "6b22fb11-0629-4c64-b1b8-be63efa293f8", Date: time.Now()},
	}

	_, _ = r.Collection.RemoveAll(nil)
	tmp := []interface{}{}
	for _, v := range votes {
		tmp = append(tmp, v)
	}
	_ = r.Collection.Insert(tmp...)
}
