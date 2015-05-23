package resource

import (
	"fmt"
	m "go-simple-rest/models"
	"gopkg.in/mgo.v2/bson"
)

type ResourcePost struct {
}

const colName string = "post"

func (r *ResourcePost) Create(post *m.Post) error {
	post.Id = bson.NewObjectId()
	err := collection(colName).Insert(post)
	return err
}

func (r *ResourcePost) List() ([]m.Post, error) {
	var posts []m.Post
	err := collection(colName).Find(nil).All(&posts)
	return posts, err
}

func (r *ResourcePost) Distinc() ([]m.Post, error) {
	var posts []m.Post
	err := collection(colName).Find(nil).Distinct("content", &posts)

	var result []string

	err = collection(colName).Find(nil).Distinct("content", &result)
	return posts, err
}

func (r *ResourcePost) GetById(id bson.ObjectId) (m.Post, error) {
	var post m.Post
	err := collection(colName).FindId(id).One(&post)
	fmt.Println(err)
	return post, err
}

func (r *ResourcePost) Update(post *m.Post) error {
	err := collection(colName).UpdateId(post.Id, post)
	return err
}

func (r *ResourcePost) Remove(id bson.ObjectId) error {
	err := collection(colName).RemoveId(id)
	return err
}
