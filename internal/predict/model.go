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
	Category    string `json:"category"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type History struct {
	Id        string `json:"id" db:"id"`
	CreatedAt string `json:"created_at" db:"created_at"`
	Result    `json:"result"`
	UserId    string `json:"user_id"`
}

type Result struct {
	Data string `json:"data"`
}
