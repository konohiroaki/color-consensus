package user

import (
	"github.com/globalsign/mgo"
	"github.com/twinj/uuid"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type User struct {
	ID string `json:"id"`
	// https://ja.wikipedia.org/wiki/ISO_3166-1
	Nationality string    `json:"nationality" bson:"nationality"`
	Gender      string    `json:"gender" bson:"gender"`
	Birth       int       `json:"birth" bson:"birth"`
	Date        time.Time `json:"date" bson:"date"`
}

var userCollection *mgo.Collection

func InitRepo(uri, db string) {
	session, _ := mgo.Dial(uri)
	c := session.DB(db).C("user")
	userCollection = c
}

func Get(id string) (User, bool) {
	var user User
	if err := userCollection.Find(bson.M{"id": id}).One(&user); err != nil {
		return User{}, false
	}
	return user, true
}

func GetList() []User {
	var userList []User
	_ = userCollection.Find(bson.M{}).All(&userList)
	if userList == nil {
		return []User{}
	}
	return userList
}

func Add(user User) string {
	user.ID = uuid.NewV4().String()
	user.Date = time.Now()
	_ = userCollection.Insert(&user)
	return user.ID
}

func InsertSampleData() {
	users := []*User{
		{ID: "00943efe-0aa5-46a4-ae5b-6ef818fc1480", Nationality: "Japan", Gender: "Male", Birth: 1985, Date: time.Now()},
		{ID: "0da04f70-dc71-4674-b47b-365c3b0805c4", Nationality: "America", Gender: "Male", Birth: 1990, Date: time.Now()},
		{ID: "20af3406-8c7e-411a-851f-31732416fa83", Nationality: "Japan", Gender: "Female", Birth: 1995, Date: time.Now()},
	}

	_, _ = userCollection.RemoveAll(nil)
	tmp := []interface{}{}
	for _, v := range users {
		tmp = append(tmp, v)
	}
	_ = userCollection.Insert(tmp...)
}
