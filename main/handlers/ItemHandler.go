package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
	"wynnpenguinapi/main/datastructures"
)

var items []datastructures.Item
var itemVersion float64

func InitializeItems() {
	jsonFile, err := os.Open("data/items.json")
	if err != nil {
		panic(err)
	}
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(jsonFile)
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var data datastructures.Itemrequest
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		panic(err)
	}
	//Iterate over all items
	for _, element := range data.Items {
		items = append(items, element)
	}
	itemVersion = data.Request.Version
	fmt.Print(len(items))
	fmt.Println(" items loaded")
}

// ItemsIndex A handler to give out all the items stored in the api.
func ItemsIndex(w http.ResponseWriter, r *http.Request) {
	itemresponse := datastructures.Itemrequest{Items: items,
		Request: datastructures.RequestStruct{Timestamp: time.Now().Unix(), Version: itemVersion}}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(itemresponse); err != nil {
		_, err := fmt.Fprint(w, "Internal Server error")
		if err != nil {
			return
		}
	}
}

// SearchItem A handler to give out a specific item or multiple items up to 10 items at a time.
// Reasoning: Requesting a whole build is now a lot easier
// Multiple Items should be seperated by a ',' the list can be in {} or not.
func SearchItem(w http.ResponseWriter, r *http.Request) {
	searches := mux.Vars(r)
	searches["name"] = strings.Replace(searches["name"], "{", "", 1)
	searches["name"] = strings.Replace(searches["name"], "}", "", 1)
	requested := strings.Split(searches["name"], ",")

	if len(requested) > 10 {
		_, err := fmt.Fprintln(w, "Too many requests")
		if err != nil {
			return
		}
		return
	}

	var result []datastructures.Item

	for _, name := range requested {
		for index, i := range items {
			if i.Name == name {
				result = append(result, items[index])
			}
		}
	}

	if len(result) <= 0 {
		_, err := fmt.Fprint(w, "empty")
		if err != nil {
			return
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	itemResponse := datastructures.Itemrequest{Items: result,
		Request: datastructures.RequestStruct{Timestamp: time.Now().Unix(), Version: itemVersion}}

	if err := json.NewEncoder(w).Encode(itemResponse); err != nil {
		panic(err)
	}

}
