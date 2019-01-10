package entities

type Invoice struct {
	Id       string   `json:"id"`
	Name     string   `json:"name"`
	Payment  string   `json:"payment"`
	Total    int      `json:"total"`
	Customer Customer `json:"customer"`
	Status   string   `json:"status"`
}
