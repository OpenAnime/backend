package constani

type PlayList struct {
	Id          string   `json:"id"`
	Name        string   `json:"title"`
	Description string   `json:"description"`
	Image       Image    `json:"image"`
	Items       []string `json:"items"`
	Slot        int      `json:"slot"`
	CreatedAt   string   `json:"created_at"`
}

type Image struct {
	Avatar string `json:"avatar"`
	Banner string `json:"banner"`
}
