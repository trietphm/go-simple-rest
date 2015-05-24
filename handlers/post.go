package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-simple-rest/models"
	"gopkg.in/mgo.v2/bson"
	"io"
	"io/ioutil"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {
	setHeaderAndEncode(w, http.StatusOK)
	var posts []models.Post
	var err error
	if posts, err = postResource.List(); err != nil {
		panic(err)
	}
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		panic(err)
	}
}

func Distinc(w http.ResponseWriter, r *http.Request) {
	setHeaderAndEncode(w, http.StatusOK)
	var posts []models.Post
	var err error
	if posts, err = postResource.Distinc(); err != nil {
		panic(err)
	}
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		panic(err)
	}
}

func Query(w http.ResponseWriter, r *http.Request) {
	setHeaderAndEncode(w, http.StatusOK)
	var posts []models.Post
	var err error
	if posts, err = postResource.Query(); err != nil {
		panic(err)
	}
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		panic(err)
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &post); err != nil {
		setHeaderAndEncode(w, 422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	} else {
		if err := postResource.Create(&post); err != nil {
			panic(err)
		}

		setHeaderAndEncode(w, http.StatusOK)
		if err := json.NewEncoder(w).Encode(post); err != nil {
			panic(err)
		}
	}

}

func GetPostById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}
	setHeaderAndEncode(w, http.StatusOK)
	oid := bson.ObjectIdHex(id)

	post, err := postResource.GetById(oid)
	if err != nil {
		panic(err)
	}
	// fmt.Fprintln(w, id, posts, post)
	if err := json.NewEncoder(w).Encode(post); err != nil {
		panic(err)
	}
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	var input models.Post
	// vars := mux.Vars(r)
	// id := vars["id"]
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 104768))
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	// fmt.Println(body)
	if err := json.Unmarshal(body, &input); err != nil {
		// fmt.Println("FAIL")
		setHeaderAndEncode(w, 422)
		panic(err)
	}

	setHeaderAndEncode(w, http.StatusOK)
	if err := postResource.Update(&input); err != nil {
		panic(err)
	}
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}
	setHeaderAndEncode(w, http.StatusOK)
	oid := bson.ObjectIdHex(id)

	if err := postResource.Remove(oid); err != nil {
		panic(err)
	}
	setHeaderAndEncode(w, http.StatusOK)

}

func setHeaderAndEncode(w http.ResponseWriter, httpStatus int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(httpStatus)
}
