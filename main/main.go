package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"wynnpenguinapi/main/handlers"
)

//TODO: Write Files
//file, _ := json.Marshal(data)
//	_ = ioutil.WriteFile("server/main/data/items.json", file, 0644)

func main() {
	fmt.Println("Initializing server...")
	//Initialize items
	handlers.InitializeItems()

	fmt.Println("Establishing connections")

	// Initialize routes of the server
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index).Methods("GET")
	router.HandleFunc("/items", handlers.ItemsIndex).Methods("GET")
	router.HandleFunc("/cache/get/itemList", handlers.ItemsIndex).Methods("GET")
	router.HandleFunc("/itemDB", handlers.ItemsIndex).Methods("GET")
	router.HandleFunc("/items/search={name}", handlers.SearchItem).Methods("GET")
	router.HandleFunc("/itemDB/search={name}", handlers.SearchItem).Methods("GET")
	router.HandleFunc("/player/search={name}", handlers.SearchPlayer).Methods("GET")
	router.HandleFunc("/playerDB/search={name}", handlers.SearchPlayer).Methods("GET")
	router.HandleFunc("/player/{name}", handlers.SearchPlayer).Methods("GET")
	router.HandleFunc("/playerDB/{name}", handlers.SearchPlayer).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Index A basic function that lists all functions
func Index(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "Welcome!")
	if err != nil {
		return
	}
}
