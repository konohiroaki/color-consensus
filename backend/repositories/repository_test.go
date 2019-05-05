package repositories

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/dbtest"
	"testing"
)

var testDBServer *dbtest.DBServer

func TestMain(m *testing.M) {
	testDBServer = &dbtest.DBServer{}
	testDBServer.SetPath("./mock_repositories/dbtest/")

	m.Run()

	testDBServer.Stop()
}

func setup() (*mgo.Database, func()) {
	session := testDBServer.Session()
	testDB := session.DB("test")

	return testDB, func() {
		session.Close()
		testDBServer.Wipe()
	}
}
