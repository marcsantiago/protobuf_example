package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"./model"
	"github.com/golang/protobuf/proto"
)

func main() {
	var newPerson model.Person

	res, err := http.Get("http://localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	err = proto.Unmarshal(data, &newPerson)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	fmt.Printf("first name is %s\n\n", newPerson.FirstName)
	// test to show that protobufer objects are json safe
	b, _ := json.MarshalIndent(newPerson, "", " ")
	fmt.Println(string(b))

	fmt.Printf("\nProto size: %d, json size: %d Protobuffer is %.2f percent the size of json\n", len(data), len(b), 100*(1/(float64(len(b))/float64(len(data)))))

}
