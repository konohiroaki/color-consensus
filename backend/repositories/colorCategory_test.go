package repositories

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestColorCategoryRepository_GetAll(t *testing.T) {
	testDB, teardown := setup()
	defer teardown(t)
	colorCategoryRepo := newColorCategoryRepository(testDB)

	name, userID := "X11", "testuser"
	_ = testDB.C("colorCategory").Insert(colorCategory{Name: name, User: userID, Date: time.Now()})

	actual := colorCategoryRepo.GetAll()

	if len(actual) == 1 {
		assert.Equal(t, name, actual[0])
	} else {
		t.Fatalf("number of documents should be exactly 1, but found %d", len(actual))
	}
}

func TestColorCategoryRepository_GetAll_Empty(t *testing.T) {
	testDB, teardown := setup()
	defer teardown(t)
	colorCategoryRepo := newColorCategoryRepository(testDB)

	actual := colorCategoryRepo.GetAll()

	assert.Len(t, actual, 0)
}

func TestColorCategoryRepository_IsPresent_True(t *testing.T) {
	testDB, teardown := setup()
	defer teardown(t)
	colorCategoryRepo := newColorCategoryRepository(testDB)

	name, userID := "X11", "testuser"
	_ = testDB.C("colorCategory").Insert(colorCategory{Name: name, User: userID, Date: time.Now()})

	actual := colorCategoryRepo.IsPresent(name)

	assert.True(t, actual)
}

func TestColorCategoryRepository_IsPresent_False(t *testing.T) {
	testDB, teardown := setup()
	defer teardown(t)
	colorCategoryRepo := newColorCategoryRepository(testDB)

	actual := colorCategoryRepo.IsPresent("X11")

	assert.False(t, actual)
}
