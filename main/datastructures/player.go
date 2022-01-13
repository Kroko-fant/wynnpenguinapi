package datastructures

import "time"

type PlayerRequest struct {
	Kind      string `json:"kind"`
	Code      int    `json:"code"`
	Timestamp int64  `json:"timestamp"`
	Version   string `json:"version"`
	Data      []struct {
		Username string `json:"username"`
		Uuid     string `json:"uuid"`
		Rank     string `json:"rank"`
		Meta     struct {
			FirstJoin time.Time `json:"firstJoin"`
			LastJoin  time.Time `json:"lastJoin"`
			Location  struct {
				Online bool   `json:"online"`
				Server string `json:"server"`
			} `json:"location"`
			Playtime int `json:"playtime"`
			Tag      struct {
				Display bool   `json:"display"`
				Value   string `json:"value"`
			} `json:"tag"`
			Veteran bool `json:"veteran"`
		} `json:"meta"`
		Classes []struct {
			Name     string `json:"name"`
			Level    int    `json:"level"`
			Dungeons struct {
				Completed int `json:"completed"`
				List      []struct {
					Name      string `json:"name"`
					Completed int    `json:"completed"`
				} `json:"list"`
			} `json:"dungeons"`
			Raids struct {
				Completed int `json:"completed"`
				List      []struct {
					Name      string `json:"name"`
					Completed int    `json:"completed"`
				} `json:"list"`
			} `json:"raids"`
			Quests struct {
				Completed int      `json:"completed"`
				List      []string `json:"list"`
			} `json:"quests"`
			ItemsIdentified int `json:"itemsIdentified"`
			MobsKilled      int `json:"mobsKilled"`
			Pvp             struct {
				Kills  int `json:"kills"`
				Deaths int `json:"deaths"`
			} `json:"pvp"`
			ChestsFound  int `json:"chestsFound"`
			BlocksWalked int `json:"blocksWalked"`
			Logins       int `json:"logins"`
			Deaths       int `json:"deaths"`
			Playtime     int `json:"playtime"`
			Gamemode     struct {
				Craftsman bool `json:"craftsman"`
				Hardcore  bool `json:"hardcore"`
				Ironman   bool `json:"ironman"`
				Hunted    bool `json:"hunted"`
			} `json:"gamemode"`
			Skills struct {
				Strength     int `json:"strength"`
				Dexterity    int `json:"dexterity"`
				Intelligence int `json:"intelligence"`
				Defence      int `json:"defence"`
				Defense      int `json:"defense"`
				Agility      int `json:"agility"`
			} `json:"skills"`
			Professions struct {
				Alchemism struct {
					Level int     `json:"level"`
					Xp    float64 `json:"xp"`
				} `json:"alchemism"`
				Armouring struct {
					Level int     `json:"level"`
					Xp    float64 `json:"xp"`
				} `json:"armouring"`
				Combat struct {
					Level int     `json:"level"`
					Xp    float64 `json:"xp"`
				} `json:"combat"`
				Cooking struct {
					Level int     `json:"level"`
					Xp    float64 `json:"xp"`
				} `json:"cooking"`
				Farming struct {
					Level int     `json:"level"`
					Xp    float64 `json:"xp"`
				} `json:"farming"`
				Fishing struct {
					Level int     `json:"level"`
					Xp    float64 `json:"xp"`
				} `json:"fishing"`
				Jeweling struct {
					Level int     `json:"level"`
					Xp    float64 `json:"xp"`
				} `json:"jeweling"`
				Mining struct {
					Level int     `json:"level"`
					Xp    float64 `json:"xp"`
				} `json:"mining"`
				Scribing struct {
					Level int     `json:"level"`
					Xp    float64 `json:"xp"`
				} `json:"scribing"`
				Tailoring struct {
					Level int     `json:"level"`
					Xp    float64 `json:"xp"`
				} `json:"tailoring"`
				Weaponsmithing struct {
					Level int     `json:"level"`
					Xp    float64 `json:"xp"`
				} `json:"weaponsmithing"`
				Woodcutting struct {
					Level int     `json:"level"`
					Xp    float64 `json:"xp"`
				} `json:"woodcutting"`
				Woodworking struct {
					Level int     `json:"level"`
					Xp    float64 `json:"xp"`
				} `json:"woodworking"`
			} `json:"professions"`
			Discoveries      int  `json:"discoveries"`
			EventsWon        int  `json:"eventsWon"`
			PreEconomyUpdate bool `json:"preEconomyUpdate"`
		} `json:"classes"`
		Guild struct {
			Name string `json:"name"`
			Rank string `json:"rank"`
		} `json:"guild"`
		Global struct {
			ChestsFound     int   `json:"chestsFound"`
			BlocksWalked    int64 `json:"blocksWalked"`
			ItemsIdentified int   `json:"itemsIdentified"`
			MobsKilled      int   `json:"mobsKilled"`
			TotalLevel      struct {
				Combat     int `json:"combat"`
				Profession int `json:"profession"`
				Combined   int `json:"combined"`
			} `json:"totalLevel"`
			Pvp struct {
				Kills  int `json:"kills"`
				Deaths int `json:"deaths"`
			} `json:"pvp"`
			Logins      int `json:"logins"`
			Deaths      int `json:"deaths"`
			Discoveries int `json:"discoveries"`
			EventsWon   int `json:"eventsWon"`
		} `json:"global"`
		Ranking struct {
			Guild  interface{} `json:"guild"`
			Player struct {
				Solo struct {
					Combat         int         `json:"combat"`
					Woodcutting    int         `json:"woodcutting"`
					Mining         interface{} `json:"mining"`
					Fishing        int         `json:"fishing"`
					Farming        int         `json:"farming"`
					Alchemism      int         `json:"alchemism"`
					Armouring      int         `json:"armouring"`
					Cooking        int         `json:"cooking"`
					Jeweling       int         `json:"jeweling"`
					Scribing       int         `json:"scribing"`
					Tailoring      int         `json:"tailoring"`
					Weaponsmithing int         `json:"weaponsmithing"`
					Woodworking    int         `json:"woodworking"`
					Profession     int         `json:"profession"`
					Overall        int         `json:"overall"`
				} `json:"solo"`
				Overall struct {
					All        int `json:"all"`
					Combat     int `json:"combat"`
					Profession int `json:"profession"`
				} `json:"overall"`
			} `json:"player"`
			Pvp interface{} `json:"pvp"`
		} `json:"ranking"`
	} `json:"data"`
}
