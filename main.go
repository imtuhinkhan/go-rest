package main

type book struct {
	ID       string  `json:"id"`
	Title    string  `json:"title"`
	Author   string  `json:"author"`
	Quantity int     `json:"quantity"`
	Price    float32 `json:"price"`
}

type author struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
