package repositories

import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"log"
	"sync"
	"time"
)

type ColorCategoryRepository interface {
	Add(name, user string) error
	GetAll() []string
}

type colorCategoryRepository struct {
	Collection *mgo.Collection
}

var (
	colorCategoryRepoInstance ColorCategoryRepository
	colorCategoryRepoOnce     sync.Once
)

func GetColorCategoryRepository(env string) ColorCategoryRepository {
	colorCategoryRepoOnce.Do(func() {
		uri, name := getDatabaseURIAndName()
		session, _ := mgo.Dial(uri)
		database := session.DB(name)
		repository := newColorCategoryRepository(database)

		if env == "development" {
			repository.insertSampleData()
		}
		colorCategoryRepoInstance = repository
	})
	return colorCategoryRepoInstance
}

func newColorCategoryRepository(database *mgo.Database) *colorCategoryRepository {
	return &colorCategoryRepository{database.C("colorCategory")}
}

type colorCategory struct {
	Name string    `bson:"name"`
	User string    `bson:"user"`
	Date time.Time `bson:"date"`
}

func (r colorCategoryRepository) Add(name, userID string) error {
	count, _ := r.Collection.Find(bson.M{"name": name}).Limit(1).Count()
	if count != 0 {
		return fmt.Errorf("the requested color category already exists")
	}
	err := r.Collection.Insert(color{
		Name: name,
		User: userID,
		Date: time.Now(),
	})
	if err != nil {
		log.Println(err)
	}
	return err
}

func (r colorCategoryRepository) GetAll() []string {
	var result []colorCategory
	err := r.Collection.
		Pipe([]bson.M{{"$project": bson.M{"_id": 0, "name": 1}}}).
		All(&result)

	if result == nil {
		if err != nil {
			log.Println(err)
		}
		return []string{}
	}

	var arrayResult = []string{}
	for _, v := range result {
		arrayResult = append(arrayResult, v.Name)
	}
	return arrayResult
}

func (r colorCategoryRepository) insertSampleData() {
	votes := []*colorCategory{
		{Name: "X11", User: "00943efe-0aa5-46a4-ae5b-6ef818fc1480", Date: time.Now()},
		{Name: "Web Color", User: "0da04f70-dc71-4674-b47b-365c3b0805c4", Date: time.Now()},
		{Name: "JIS慣用色名", User: "0da04f70-dc71-4674-b47b-365c3b0805c4", Date: time.Now()},
		{Name: "Japanese", User: "0da04f70-dc71-4674-b47b-365c3b0805c4", Date: time.Now()},
		{Name: "English", User: "0da04f70-dc71-4674-b47b-365c3b0805c4", Date: time.Now()},
	}

	_, _ = r.Collection.RemoveAll(nil)
	tmp := []interface{}{}
	for _, v := range votes {
		tmp = append(tmp, v)
	}
	_ = r.Collection.Insert(tmp...)
}
