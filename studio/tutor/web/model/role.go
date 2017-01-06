package model



type Role struct {
	Basic
	MenuActions []MenuAction `bson: "menuActions"`
}
