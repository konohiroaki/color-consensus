package repository

import (
	"github.com/globalsign/mgo"
	"github.com/konohiroaki/color-consensus/backend/models"
	"time"
)

type userRepository struct {
	Collection *mgo.Collection
}

var userRepo *userRepository

func InitUserRepo() {
	// TODO: use env variable
	uri := "mongodb://localhost:27017/"
	session, _ := mgo.Dial(uri)
	c := session.DB("cc").C("user")
	userRepo = &userRepository{Collection: c}
}

func GetUserRepo() *userRepository {
	return userRepo
}

func (u *userRepository) InsertSampleUser() {
	_ = u.Collection.Insert(&models.User{ID: "foo", Nationality: "Japan", Gender: "male", Birth: 1990, Date: time.Now()})
}

func InsertSampleUser() {
	users := []*models.User{
		{ID: "0da04f70-dc71-4674-b47b-365c3b0805c4", Nationality: "Japan", Gender: "Male", Birth: 1990, Date: time.Now()},
		{ID: "20af3406-8c7e-411a-851f-31732416fa83", Nationality: "Japan", Gender: "Male", Birth: 1991, Date: time.Now()},
	}

	c := userRepo.Collection
	_, _ = c.RemoveAll(nil)
	tmp := []interface{}{}
	for _, v := range users {
		tmp = append(tmp, v)
	}
	_ = c.Insert(tmp...)
}
