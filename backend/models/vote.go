package models

import "time"

type ColorVote struct {
	Language string    `json:"-"`
	Color    string    `json:"-"`
	User     string    `json:"-"`
	Date     time.Time `json:"date"`
	//FIXME: validate not working.
	Colors []string `json:"colors" validate:"dive,hexcolor"`
}