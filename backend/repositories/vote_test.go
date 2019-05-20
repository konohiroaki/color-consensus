package repositories

import (
	"github.com/globalsign/mgo/bson"
	"github.com/stretchr/testify/assert"
	"github.com/xeonx/timeago"
	"math"
	"testing"
	"time"
)

func TestVoteRepository_Add_Insert(t *testing.T) {
	testDB, teardown := setup()
	defer teardown(t)
	voteRepo := newVoteRepository(testDB)

	category, name, colors, userID := "HTML Color", "Red", []string{"#ff0000"}, "testuser"
	voteRepo.Add(category, name, colors, userID)

	var actual []vote
	_ = testDB.C("vote").Find(bson.M{}).All(&actual)

	if len(actual) == 1 {
		assert.Equal(t, category, actual[0].Category)
		assert.Equal(t, name, actual[0].Name)
		assert.Equal(t, colors, actual[0].Colors)
		assert.True(t, actual[0].Date.After(time.Now().Add(-10*time.Second)) && actual[0].Date.Before(time.Now()))
		assert.Equal(t, actual[0].User, userID)
	} else {
		t.Fatalf("number of documents should be exactly 1, but found %d", len(actual))
	}
}

func TestVoteRepository_Add_Update(t *testing.T) {
	testDB, teardown := setup()
	defer teardown(t)
	voteRepo := newVoteRepository(testDB)

	category, name, colors, userID := "HTML Color", "Red", []string{"#ff0000"}, "testuser"
	_ = testDB.C("vote").Insert(vote{Category: category, Name: name, Colors: []string{"#ff1010"}, Date: time.Now().Add(-timeago.Year), User: userID})
	voteRepo.Add(category, name, colors, userID)

	var actual []vote
	_ = testDB.C("vote").Find(bson.M{}).All(&actual)

	if len(actual) == 1 {
		assert.Equal(t, category, actual[0].Category)
		assert.Equal(t, name, actual[0].Name)
		assert.Equal(t, colors, actual[0].Colors)
		assert.True(t, actual[0].Date.After(time.Now().Add(-10*time.Second)) && actual[0].Date.Before(time.Now()))
		assert.Equal(t, actual[0].User, userID)
	} else {
		t.Fatalf("number of documents should be exactly 1, but found %d", len(actual))
	}
}

func TestVoteRepository_Get(t *testing.T) {
	testDB, teardown := setup()
	defer teardown(t)
	voteRepo := newVoteRepository(testDB)

	nationality, birth, gender := "foo", 1970, "bar"
	category, name, colors, userID := "HTML Color", "Red", []string{"#ff0000"}, "testuser"
	_ = testDB.C("user").Insert(user{ID: userID, Nationality: nationality, Birth: birth, Gender: gender, Date: time.Now()})
	_ = testDB.C("vote").Insert(vote{Category: category, Name: name, Colors: colors, Date: time.Now(), User: userID})

	actual := voteRepo.Get(category, name, []string{"voter.nationality", "voter.ageGroup", "voter.gender", "category", "name", "colors"})

	if len(actual) == 1 {
		actualVoter := actual[0]["voter"].(map[string]interface{})
		assert.Equal(t, nationality, actualVoter["nationality"])
		expectedAgeGroup := math.Floor(float64(time.Now().Year()-birth)/10) * 10
		assert.Equal(t, expectedAgeGroup, actualVoter["ageGroup"].(float64))
		assert.Equal(t, gender, actualVoter["gender"])
		assert.Equal(t, category, actual[0]["category"])
		assert.Equal(t, name, actual[0]["name"])
		assert.Equal(t, colors[0], actual[0]["colors"].([]interface{})[0])
		assert.Equal(t, nil, actual[0]["date"])
	} else {
		t.Fatalf("number of documents should be exactly 1, but found %d", len(actual))
	}
}

func TestVoteRepository_Get_Empty(t *testing.T) {
	testDB, teardown := setup()
	defer teardown(t)
	voteRepo := newVoteRepository(testDB)

	category, name := "HTML Color", "Red"

	actual := voteRepo.Get(category, name, []string{"voter.nationality", "voter.ageGroup", "voter.gender", "category", "name", "colors"})

	assert.Len(t, actual, 0)
}

func TestVoteRepository_RemoveByUser(t *testing.T) {
	testDB, teardown := setup()
	defer teardown(t)
	voteRepo := newVoteRepository(testDB)

	category, name, colors, userID, userID2 := "HTML Color", "Red", []string{"#ff0000"}, "testuser", "testuser2"
	_ = testDB.C("vote").Insert(vote{Category: category, Name: name, Colors: colors, Date: time.Now(), User: userID})
	_ = testDB.C("vote").Insert(vote{Category: category, Name: name, Colors: colors, Date: time.Now(), User: userID2})
	voteRepo.RemoveByUser(userID)

	var actual []vote
	_ = testDB.C("vote").Find(bson.M{}).All(&actual)

	if len(actual) == 1 {
		assert.Equal(t, userID2, actual[0].User)
	} else {
		t.Fatalf("number of documents should be exactly 1, but found %d", len(actual))
	}
}
