package models

import "time"

/*
 * In DB, associate this data with cookie data. Also, generate a key to reuse that user data when cookie is gone.
 * If user have data, just go to the input screen.
 * User should be able to list input history in some way.
 */
type User struct {
	ID string `json:"id"`
	// https://ja.wikipedia.org/wiki/ISO_3166-1
	Nationality string    `json:"nationality"`
	Gender      string    `json:"gender"`
	Birth       uint      `json:"birth"`
	Date        time.Time `json:"date"`
}
