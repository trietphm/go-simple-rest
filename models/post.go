package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Post struct {
	Id      bson.ObjectId `bson:"_id" json:"id"`
	Time    time.Time     `bson:"time" json:"time"`
	Content string        `bson:"content" json:"content"`
}
