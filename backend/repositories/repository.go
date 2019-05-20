package repositories

import (
	"github.com/konohiroaki/color-consensus/backend/config"
	"os"
	"strings"
)

func getDatabaseURIAndName() (string, string) {
	uri := os.Getenv("MONGODB_URI") // provided by mLab add-on
	db := uri[strings.LastIndex(uri, "/")+1:]
	if uri == "" {
		uri = config.GetConfig().Get("mongo.url").(string)
		db = "cc"
	}
	return uri, db
}

type DuplicateError struct {
	message string
}

func NewDuplicateError(message string) *DuplicateError {
	return &DuplicateError{message}
}

func (e DuplicateError) Error() string {
	return e.message
}
