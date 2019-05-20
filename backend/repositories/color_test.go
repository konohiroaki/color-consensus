package repositories

import (
	"github.com/globalsign/mgo/bson"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestColorRepository_Add(t *testing.T) {
	testDB, teardown := setup()
	defer teardown(t)
	colorRepo := newColorRepository(testDB)

	category, name, code, userID := "X11", "Red", "#ff0000", "testuser"
	_ = colorRepo.Add(category, name, code, userID)

	var actual []color
	_ = testDB.C("color").Find(bson.M{}).All(&actual)

	if len(actual) == 1 {
		assert.Equal(t, category, actual[0].Category)
		assert.Equal(t, name, actual[0].Name)
		assert.Equal(t, code, actual[0].Code)
		assert.Equal(t, userID, actual[0].User)
		assert.True(t, actual[0].Date.After(time.Now().Add(-10*time.Second)) && actual[0].Date.Before(time.Now()))
	} else {
		t.Fatalf("number of documents should be exactly 1, but found %d", len(actual))
	}
}

func TestColorRepository_Add_SameColorPresent(t *testing.T) {
	testDB, teardown := setup()
	defer teardown(t)
	colorRepo := newColorRepository(testDB)

	category, name, code, userID := "X11", "Red", "#ff0000", "testuser"
	_ = testDB.C("color").Insert(color{Category: category, Name: name, Code: code, User: userID, Date: time.Now()})

	actual := colorRepo.Add(category, name, "#ff0011", "anotheruser")

	assert.Equal(t, "the requested color already exists", actual.Error())
}

func TestColorRepository_GetAll(t *testing.T) {
	testDB, teardown := setup()
	defer teardown(t)
	colorRepo := newColorRepository(testDB)

	category, name, code, userID := "X11", "red", "#ff0000", "testuser"
	_ = testDB.C("color").Insert(color{Category: category, Name: name, Code: code, User: userID, Date: time.Now()})

	actual := colorRepo.GetAll([]string{"category", "name", "code"})

	if len(actual) == 1 {
		assert.Equal(t, category, actual[0]["category"])
		assert.Equal(t, name, actual[0]["name"])
		assert.Equal(t, code, actual[0]["code"])
		assert.Equal(t, nil, actual[0]["user"])
		assert.Equal(t, nil, actual[0]["date"])
	} else {
		t.Fatalf("number of documents should be exactly 1, but found %d", len(actual))
	}
}

func TestColorRepository_GetAll_Empty(t *testing.T) {
	testDB, teardown := setup()
	defer teardown(t)
	colorRepo := newColorRepository(testDB)

	actual := colorRepo.GetAll([]string{"category", "name", "code"})

	assert.Len(t, actual, 0)
}
