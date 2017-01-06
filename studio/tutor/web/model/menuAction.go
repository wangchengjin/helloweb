package model

type ActionType byte

type MenuAction struct {
	Basic
	Menu       Menu
	ActionType ActionType `bson:"actionType"` //
}
