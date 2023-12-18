package foods

type Food struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Image       string `json:"image" db:"image"`
	Category    string `json:"category" db:"category"`
}
