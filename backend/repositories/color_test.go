package repositories

import (
	"github.com/globalsign/mgo/bson"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestColorRepository_Add(t *testing.T) {
	testDB, teardown := setup()
	defer teardown()
	colorRepo := newColorRepository(testDB)

	lang, name, code, userID := "en", "red", "#ff0000", "testuser"
	colorRepo.Add(lang, name, code, userID)

	var actual []map[string]interface{}
	_ = testDB.C("color").Find(bson.M{}).All(&actual)

	assert.Equal(t, actual[0]["lang"], lang)
	assert.Equal(t, actual[0]["name"], name)
	assert.Equal(t, actual[0]["code"], code)
	assert.Equal(t, actual[0]["user"], userID)
	date, now := actual[0]["date"].(time.Time), time.Now()
	assert.True(t, date.After(now.Add(-10*time.Second)) && date.Before(now))
}

func TestColorRepository_GetAll(t *testing.T) {
	testDB, teardown := setup()
	defer teardown()
	colorRepo := newColorRepository(testDB)

	lang, name, code, userID := "en", "red", "#ff0000", "testuser"
	_ = testDB.C("color").Insert(color{Lang: lang, Name: name, Code: code, User: userID, Date: time.Now()})

	actual := colorRepo.GetAll([]string{"lang", "name", "code"})

	assert.Equal(t, actual[0]["lang"], lang)
	assert.Equal(t, actual[0]["name"], name)
	assert.Equal(t, actual[0]["code"], code)
	assert.Equal(t, actual[0]["date"], nil)
}
