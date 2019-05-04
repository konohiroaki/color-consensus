package repositories

import (
	"github.com/globalsign/mgo/bson"
	"github.com/globalsign/mgo/dbtest"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestColorRepository_Add(t *testing.T) {
	dbServer := dbtest.DBServer{}
	dbServer.SetPath("./mock_repositories/dbtest/")
	session := dbServer.Session()
	database := session.DB("test")
	colorRepo := newColorRepository(database)

	lang, name, code, userID := "en", "red", "#ff0000", "testuser"
	colorRepo.Add(lang, name, code, userID)

	var actual []map[string]string
	_ = database.C("color").Find(bson.M{}).All(&actual)
	assert.Equal(t, actual[0]["lang"], lang)
	assert.Equal(t, actual[0]["name"], name)
	assert.Equal(t, actual[0]["code"], code)
	assert.Equal(t, actual[0]["user"], userID)

	session.Close()

	dbServer.Wipe()
	dbServer.Stop()
}
