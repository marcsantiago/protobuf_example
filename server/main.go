package main

import (
	"log"
	"net/http"

	"../model"
	"github.com/golang/protobuf/proto"
)

func getPerson(w http.ResponseWriter, r *http.Request) {
	person := &model.Person{
		FirstName: "Marc",
		LastName:  "Santiago",
		Age:       26,
		Parents:   []*model.Person_Parent{&model.Person_Parent{FirstName: "Marc", LastName: "Santiago"}},
	}

	data, err := proto.Marshal(person)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	w.Write(data)
	return
}

func main() {
	http.HandleFunc("/", getPerson)
	http.ListenAndServe(":8000", nil)
}
