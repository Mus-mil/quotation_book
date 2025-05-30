package models

type QuoteBook struct {
	Author string `json:"author"`
	Quote  string `json:"quote"`
}

type QuoteBookID struct {
	ID     int    `json:"id"`
	Author string `json:"author"`
	Quote  string `json:"quote"`
}
