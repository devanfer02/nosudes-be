package domain

type GmapsRef struct {
	Results []MapsDetail `json:"results"`
}

type MapsDetail struct {
	Name         string   `json:"name"`
	Address      string   `json:"formatted_address"`
	Rating       float64  `json:"rating"`
	TotalRatings int64    `json:"user_ratings_total"`
	GeometryInfo Geometry `json:"geometry"`
}

type Geometry struct {
	Loc Location `json:"location"`
}

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
