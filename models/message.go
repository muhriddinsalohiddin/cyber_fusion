package models

type Message struct {
	Id         string `json:"id"`
	SenderId   string `json:"sender_id"`
	ReceiverId string `json:"receiver_id"`
	Body       string `json:"body"`
	CreatedAt  string `json:"created_at"`
}

type List struct {
	Messages []*Message `json:"messages"`
	Cout     int        `json:"count"`
}
