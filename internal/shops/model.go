package shops

type Shop struct {
	Id        int     `json:"id" db:"db"`
	Name      string  `json:"name" db:"name"`
	Location  string  `json:"location" db:"location"`
	GmapsLink string  `json:"gmaps_link" db:"gmaps_link"`
	Latitude  float64 `json:"latitude" db:"latitude"`
	Longitude float64 `json:"longitude" db:"longitude"`
	Image     string  `json:"image" db:"image"`
}
