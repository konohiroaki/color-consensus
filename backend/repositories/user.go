package repositories

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/twinj/uuid"
	"sync"
	"time"
)

type UserRepository interface {
	IsPresent(id string) bool
	Add(nationality string, birth int, gender string) string
	Remove(id string) error
}

type userRepository struct {
	Collection *mgo.Collection
}

var (
	userRepoInstance UserRepository
	userRepoOnce     sync.Once
)

func GetUserRepository(env string) UserRepository {
	userRepoOnce.Do(func() {
		uri, name := getDatabaseURIAndName()
		session, _ := mgo.Dial(uri)
		database := session.DB(name)
		repository := newUserRepository(database)

		if env == "development" {
			repository.insertSampleData()
		}
		userRepoInstance = repository
	})
	return userRepoInstance
}

func newUserRepository(database *mgo.Database) *userRepository {
	return &userRepository{database.C("user")}
}

type user struct {
	ID string `json:"id"`
	// https://ja.wikipedia.org/wiki/ISO_3166-1
	Nationality string    `bson:"nationality"`
	Birth       int       `bson:"birth"`
	Gender      string    `bson:"gender"`
	Date        time.Time `bson:"date"`
}

func (r userRepository) IsPresent(id string) bool {
	count, _ := r.Collection.Find(bson.M{"id": id}).Limit(1).Count()
	return count > 0
}

func (r userRepository) Add(nationality string, birth int, gender string) string {
	user := user{
		ID:          uuid.NewV4().String(),
		Nationality: nationality,
		Birth:       birth,
		Gender:      gender,
		Date:        time.Now(),
	}
	_ = r.Collection.Insert(&user)
	return user.ID
}

func (r userRepository) Remove(id string) error {
	return r.Collection.Remove(bson.M{"id": id})
}

func (r userRepository) insertSampleData() {
	users := []*user{
		{ID: "00943efe-0aa5-46a4-ae5b-6ef818fc1480", Nationality: "JP", Birth: 1985, Gender: "Male", Date: time.Now()},
		{ID: "0da04f70-dc71-4674-b47b-365c3b0805c4", Nationality: "US", Birth: 1990, Gender: "Male", Date: time.Now()},
		{ID: "20af3406-8c7e-411a-851f-31732416fa83", Nationality: "JP", Birth: 1995, Gender: "Female", Date: time.Now()},
	}

	_, _ = r.Collection.RemoveAll(nil)
	tmp := []interface{}{}
	for _, v := range users {
		tmp = append(tmp, v)
	}
	_ = r.Collection.Insert(tmp...)
}
