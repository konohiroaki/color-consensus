package repositories

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/dbtest"
	"os"
	"testing"
)

var testDBServer *dbtest.DBServer

func TestMain(m *testing.M) {
	testDBServer = &dbtest.DBServer{}
	testDBServer.SetPath("./mock_repositories/dbtest/")

	code := m.Run()

	testDBServer.Stop()
	os.Exit(code)
}

func setup() (*mgo.Database, func(t *testing.T)) {
	session := testDBServer.Session()
	testDB := session.DB("test")

	return testDB, func(t *testing.T) {
		if r := recover(); r != nil {
			t.Error(r)
		}
		session.Close()
		testDBServer.Wipe()
	}
}
