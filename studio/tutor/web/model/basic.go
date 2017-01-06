package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Basic struct {
	ID         bson.ObjectId `bson:"_id,omitempty`
	Name       string `bson:"name"`
	CreateUser User `bson:"createUser"`
	CreateTime time.Time `bson:"createTime"`
	UpdateUser User `bson:"updateUser"`
	UpdateTime time.Time `bson:"updateTime"`
	Memo       string `bson:"memo"`
}
