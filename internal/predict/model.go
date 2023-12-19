package predict

type Predict struct {
	Id          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Category    string `json:"category" db:"category"`
	CreatedAt   string `json:"created_at" db:"created_at"`
	Description string `json:"description" db:"description"`
	Image       string `json:"image" db:"image"`
}

type PredictResponse struct {
	Category string `json:"category"`
}
