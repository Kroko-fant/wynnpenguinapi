package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
	"wynnpenguinapi/main/datastructures"
)

//Constants to paths in the api
const (
	PlayerPath1 = "https://api.wynncraft.com/v2/player/"
	PlayerPath2 = "/stats/"
)

// RequestPlayer A function to request data of a player
func RequestPlayer(player string) datastructures.PlayerRequest {
	url := PlayerPath1 + player + PlayerPath2
	client := http.Client{Timeout: time.Second * 2}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	//Note: that client is not true at all but idc
	req.Header.Set("User-Agent", "Mozilla/1.00 (compatible; MSIE 1.0; Windows 3.1)")

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(err)
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	playerRequest := datastructures.PlayerRequest{}
	jsonErr := json.Unmarshal(body, &playerRequest)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return playerRequest
}

// SearchPlayer Searches one or multiple players
func SearchPlayer(w http.ResponseWriter, r *http.Request) {
	searches := mux.Vars(r)
	searches["name"] = strings.Replace(searches["name"], "{", "", 1)
	searches["name"] = strings.Replace(searches["name"], "}", "", 1)
	requested := searches["name"]
	w.Header().Set("Content-Type", "application/json")
	playerResponse := RequestPlayer(requested)

	if err := json.NewEncoder(w).Encode(playerResponse); err != nil {
		panic(err)
	}

}
