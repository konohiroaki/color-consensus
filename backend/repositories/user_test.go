package repositories

import (
	"github.com/globalsign/mgo/bson"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUserRepository_IsPresent_Present(t *testing.T) {
	testDB, teardown := setup()
	defer teardown(t)
	userRepo := newUserRepository(testDB)

	userID := "testuser"
	_ = testDB.C("user").Insert(user{ID: userID, Nationality: "foo", Birth: 1970, Gender: "bar", Date: time.Now()})

	actual := userRepo.IsPresent(userID)

	assert.True(t, actual)
}

func TestUserRepository_IsPresent_Absent(t *testing.T) {
	testDB, teardown := setup()
	defer teardown(t)
	userRepo := newUserRepository(testDB)

	userID := "testuser"

	actual := userRepo.IsPresent(userID)

	assert.False(t, actual)
}

func TestUserRepository_Add(t *testing.T) {
	testDB, teardown := setup()
	defer teardown(t)
	userRepo := newUserRepository(testDB)

	nationality, birth, gender := "foo", 1900, "bar"
	userRepo.Add(nationality, birth, gender)

	var actual []user
	_ = testDB.C("user").Find(bson.M{}).All(&actual)

	if len(actual) == 1 {
		assert.NotEmpty(t, actual[0].ID)
		assert.Equal(t, nationality, actual[0].Nationality)
		assert.Equal(t, birth, actual[0].Birth)
		assert.Equal(t, gender, actual[0].Gender)
		assert.True(t, actual[0].Date.After(time.Now().Add(-10*time.Second)) && actual[0].Date.Before(time.Now()))
	} else {
		t.Fatalf("number of documents should be exactly 1, but found %d", len(actual))
	}
}

func TestUserRepository_Remove(t *testing.T) {
	testDB, teardown := setup()
	defer teardown(t)
	userRepo := newUserRepository(testDB)

	userID := "testuser"
	_ = testDB.C("user").Insert(user{ID: userID, Nationality: "foo", Birth: 1970, Gender: "bar", Date: time.Now()})
	_ = userRepo.Remove(userID)

	var actual []user
	_ = testDB.C("user").Find(bson.M{}).All(&actual)

	assert.Len(t, actual, 0)
}
