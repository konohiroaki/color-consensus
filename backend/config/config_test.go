package config

import "testing"

func assertEquals(t *testing.T, expected, actual string) {
	if actual != expected {
		t.Fatalf("Expected: %s, Actual %s", expected, actual)
	}
}

func TestGetConfigHasPort(t *testing.T) {
	Init("./")
	assertEquals(t, "5000", GetConfig().Get("port").(string))
}

func TestGetConfigHasMongoURL(t *testing.T) {
	Init("./")
	assertEquals(t, "mongodb://localhost:27017/", GetConfig().Get("mongo.url").(string))
}
