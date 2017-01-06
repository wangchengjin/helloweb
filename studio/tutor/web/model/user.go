package model

import (
	"time"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fastweb/studio/tutor/web/dbaccess"
)

type Sex int

const (
	_ Sex = iota
	Male
	Female
)

type User struct {
	ID         bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name       string `bson:"name,omitempty" json:"name,omitempty"`
	//CreateUser User `bson:"createUser,omitempty"`
	CreateTime time.Time `bson:"createTime,omitempty" json:"createTime,omitempty"`
	//UpdateUser User `bson:"updateUser,omitempty"`
	UpdateTime time.Time `bson:"updateTime,omitempty" json:"updateTime,omitempty"`
	Memo       string `bson:"memo,omitempty" json:"memo,omitempty"`
	LoginName  string `bson:"loginName,omitempty" json:"loginName,omitempty"`
	Password   []byte `bson:"password,omitempty" json:"password,omitempty"`
	Sex        Sex `bson:"sex,omitempty" json:"sex,omitempty"`
	Birthday   time.Time `bson:"birthday,omitempty" json:"birthday,omitempty"`
	Roles      []Role `bson:"roles,omitempty" json:"birthday,omitempty"`
	//Departments [] `bson:"departments"`
}

func AddUser(user User) string {
	user.ID = bson.NewObjectId()
	user.CreateTime = time.Now()
	user.UpdateTime = user.CreateTime
	query := func(c *mgo.Collection) error {
		return c.Insert(user)
	}
	if err := dbaccess.WitchCollection("user", query); err != nil {
		return "false"
	}
	return user.ID.Hex()
}

/**
 * 获取一条记录通过objectid
 */
func GetUserById(id string) *User {
	objId := bson.ObjectIdHex(id)
	user := new(User)
	query := func(c *mgo.Collection) error {
		return c.FindId(objId).One(&user)
	}
	dbaccess.WitchCollection("user", query)
	return user
}

//获取所有的User数据
func GetAllUser() []User {
	var users []User
	query := func(c *mgo.Collection) error {
		return c.Find(nil).All(&users)
	}

	if err := dbaccess.WitchCollection("user", query); err != nil {
		return users
	}
	return users
}


//更新person数据
func UpdateUser(user *User) bool {

	query := bson.M{"_id": user.ID}
	user.UpdateTime = time.Now()
	//change := bson.M{"$set":bson.M{"updatetime":time.Now()}}
	change := bson.M{"$set":user}
	println(change)
	/*if err != nil {
		log.Fatalln(err)
		return false
	}*/
	exop := func(c *mgo.Collection) error {
		println()
		return c.Update(query, bson.M{"$set":change})
	}
	err := dbaccess.WitchCollection("user", exop)
	return err == nil
}

/**
 * 执行查询，此方法可拆分做为公共方法
 * [SearchPerson description]
 * @param {[type]} collectionName string [description]
 * @param {[type]} query          bson.M [description]
 * @param {[type]} sort           bson.M [description]
 * @param {[type]} fields         bson.M [description]
 * @param {[type]} skip           int    [description]
 * @param {[type]} limit          int)   (results      []interface{}, err error [description]
 */
func SearchPerson(collectionName string, query bson.M, sort string, fields bson.M, skip int, limit int) (results []interface{}, err error) {
	exop := func(c *mgo.Collection) error {
		return c.Find(query).Sort(sort).Select(fields).Skip(skip).Limit(limit).All(&results)
	}
	err = dbaccess.WitchCollection(collectionName, exop)
	return
}