package user

import (
	"github.com/globalsign/mgo"
	"github.com/konohiroaki/color-consensus/backend/config"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var userCollection *mgo.Collection

func InitRepo() {
	uri := config.GetConfig().Get("mongo.url").(string)
	session, _ := mgo.Dial(uri)
	c := session.DB("cc").C("user")
	userCollection = c
}

func GetList() []User {
	var userList []User
	_ = userCollection.Find(bson.M{}).All(&userList)
	return userList
}

func InsertSampleData() {
	users := []*User{
		{ID: "0da04f70-dc71-4674-b47b-365c3b0805c4", Nationality: "Japan", Gender: "Male", Birth: 1990, Date: time.Now()},
		{ID: "20af3406-8c7e-411a-851f-31732416fa83", Nationality: "Japan", Gender: "Male", Birth: 1991, Date: time.Now()},
	}

	_, _ = userCollection.RemoveAll(nil)
	tmp := []interface{}{}
	for _, v := range users {
		tmp = append(tmp, v)
	}
	_ = userCollection.Insert(tmp...)
}
