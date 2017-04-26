package store

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"github.com/sandangel/go-recipes/chap8/chap8-7/model"
	"gopkg.in/mgo.v2/bson"
	"os/user"
)

var mgoSession *mgo.Session

func createDBSession() {
	var err error
	mgoSession, err = mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:   []string{"127.0.0.1"},
		Timeout: 60 * time.Second,
	})
	if err != nil {
		log.Fatalf("[createDbSession]:%s\n", err)
	}
}

func init() {
	createDBSession()
}

type MongoUserStore struct {
}

func (store MongoUserStore) GetUsers(user model.User) []model.User {
	session := mgoSession.Copy()
	defer session.Close()
	userCol := session.DB("userdb").C("users")
	var users []model.User
	iter := userCol.Find(nil).Iter()
	result := model.User{}
	for iter.Next(&result) {
		users = append(users, result)
	}
	return users
}

func (store MongoUserStore) AddUser(user model.User) error {
	session := mgoSession.Copy()
	defer session.Close()
	userCol := session.DB("userdb").C("users")
	var existUser model.User
	err := userCol.FindId(bson.M{"email": user.Email}).One(&existUser)
	if err != nil {
		if err == mgo.ErrNotFound{
			log.Fatalf("Email is unique: %v", err)
		}
	}
	if (model.User{}) != existUser {
		return model.ErrorEmailExists
	}
	err = userCol.Insert(user)
	return err
}
