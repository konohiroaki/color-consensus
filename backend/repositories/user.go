package repositories

import (
	"fmt"
	"github.com/twinj/uuid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type UserRepository interface {
	Add(nationality, gender string, birth int) string
	IsPresent(id string) bool
}

type userRepository struct {
	Collection *mgo.Collection
}

func NewUserRepository(uri, db, env string) UserRepository {
	session, _ := mgo.Dial(uri)
	collection := session.DB(db).C("user")
	repository := &userRepository{collection}

	if env == "development" {
		fmt.Println("detected development mode. inserting sample user data.")
		repository.insertSampleData()
	}

	return repository
}

type user struct {
	ID string `json:"id"`
	// https://ja.wikipedia.org/wiki/ISO_3166-1
	Nationality string    `json:"nationality" bson:"nationality"`
	Gender      string    `json:"gender" bson:"gender"`
	Birth       int       `json:"birth" bson:"birth"`
	Date        time.Time `json:"date" bson:"date"`
}

func (r userRepository) Add(nationality, gender string, birth int) string {
	user := user{
		ID:          uuid.NewV4().String(),
		Nationality: nationality,
		Gender:      gender,
		Birth:       birth,
		Date:        time.Now(),
	}
	_ = r.Collection.Insert(&user)
	return user.ID
}

func (r userRepository) IsPresent(id string) bool {
	count, _ := r.Collection.Find(bson.M{"id": id}).Count()
	return count > 0
}

func (r userRepository) insertSampleData() {
	users := []*user{
		{ID: "00943efe-0aa5-46a4-ae5b-6ef818fc1480", Nationality: "Japan", Gender: "Male", Birth: 1985, Date: time.Now()},
		{ID: "0da04f70-dc71-4674-b47b-365c3b0805c4", Nationality: "America", Gender: "Male", Birth: 1990, Date: time.Now()},
		{ID: "20af3406-8c7e-411a-851f-31732416fa83", Nationality: "Japan", Gender: "Female", Birth: 1995, Date: time.Now()},
		{ID: "testuser", Nationality: "XXX", Gender: "XXX", Birth: 1, Date: time.Now()},
	}

	_, _ = r.Collection.RemoveAll(nil)
	tmp := []interface{}{}
	for _, v := range users {
		tmp = append(tmp, v)
	}
	_ = r.Collection.Insert(tmp...)
}
