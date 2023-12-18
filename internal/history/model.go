package history

type History struct {
	Id     string `json:"id" db:"id"`
	Result string `json:"result" db:"result"`
}
