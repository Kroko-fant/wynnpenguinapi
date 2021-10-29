package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"wynnpenguinapi/main/item"
)

var items []item.Item
var version float32

//TODO: Write Files
//file, _ := json.Marshal(data)
//	_ = ioutil.WriteFile("server/main/data/api.json", file, 0644)

func main() {
	fmt.Println("Initializing server...")
	//
	jsonFile, err := os.Open("data/api.json")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(jsonFile)
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var data item.DictWithItemList
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		fmt.Println(err)
	}
	//Iterate over all items
	for _, element := range data.Items {
		items = append(items, element)
	}
	fmt.Print("Data loaded ")
	fmt.Print(len(items))
	fmt.Println(" items loaded")
	version = 1.0

	fmt.Println("Establishing connections")

	// Initialize routes of the server
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index).Methods("GET")
	router.HandleFunc("/items", ItemsIndex).Methods("GET")
	router.HandleFunc("/cache/get/itemList", ItemsIndex).Methods("GET")
	router.HandleFunc("/itemDB", ItemsIndex).Methods("GET")
	router.HandleFunc("/items/search={name}", SearchItem).Methods("GET")
	router.HandleFunc("/itemDB/search={name}", SearchItem).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

// Index A basic function that lists all functions
func Index(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "Welcome!")
	if err != nil {
		return
	}
}

// ItemsIndex A handler to give out all the items stored in the api.
func ItemsIndex(w http.ResponseWriter, r *http.Request) {
	itemresponse := item.DictWithItemList{Items: items,
		Request: item.Request{Timestamp: time.Now().Unix(), Version: version}}
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

	var result item.List

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
	itemresponse := item.DictWithItemList{Items: result,
		Request: item.Request{Timestamp: time.Now().Unix(), Version: version}}

	if err := json.NewEncoder(w).Encode(itemresponse); err != nil {
		panic(err)
	}

}
