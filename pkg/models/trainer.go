package models

type Trainer struct {
	Name         string `json:"name"`
	Age          int    `json:"age"`
	OriginRegion string `json:"origin_region"`
	OriginCity   string `json:"origin_city"`
	BestBuddy    string `json:"best_buddy"`
	Team         string `json:"team"`
}
