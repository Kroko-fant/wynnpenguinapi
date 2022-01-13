package datastructures

import "time"

type Guild struct {
	Name    string `json:"name"`
	Prefix  string `json:"prefix"`
	Members []struct {
		Name           string    `json:"name"`
		Uuid           string    `json:"uuid"`
		Rank           string    `json:"rank"`
		Contributed    int64     `json:"contributed"`
		Joined         time.Time `json:"joined"`
		JoinedFriendly string    `json:"joinedFriendly"`
	} `json:"members"`
	Xp              float64   `json:"xp"`
	Level           int       `json:"level"`
	Created         time.Time `json:"created"`
	CreatedFriendly string    `json:"createdFriendly"`
	Territories     int       `json:"territories"`
	Banner          struct {
		Base      string `json:"base"`
		Tier      int    `json:"tier"`
		Structure string `json:"structure"`
		Layers    []struct {
			Colour  string `json:"colour"`
			Pattern string `json:"pattern"`
		} `json:"layers"`
	} `json:"banner"`
	Request struct {
		Timestamp int `json:"timestamp"`
		Version   int `json:"version"`
	} `json:"request"`
}
