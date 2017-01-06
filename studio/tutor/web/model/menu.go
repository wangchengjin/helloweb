package model

type Menu struct {
	Basic
	Url  string `bson:"url"`
	Icon string `bson:"icon"`
}
