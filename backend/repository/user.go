package repository

import (
	"github.com/globalsign/mgo"
	"github.com/konohiroaki/color-consensus/backend/models"
	"sync"
	"time"
)

type userRepository struct {
	Collection *mgo.Collection
}

var userRepo *userRepository
var userOnce sync.Once

func GetUserRepo() *userRepository {
	userOnce.Do(func() {
		// TODO: use env variable
		uri := "mongodb://localhost:27017/"
		session, _ := mgo.Dial(uri)
		c := session.DB("cc").C("user")
		userRepo = &userRepository{Collection: c}
	})
	return userRepo
}

func (u *userRepository) InsertSampleUser() {
	_ = u.Collection.Insert(&models.User{ID: "foo", Nationality: "Japan", Gender: "male", Birth: 1990, Date: time.Now()})
}
