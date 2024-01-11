package models

type Message struct {
	Id         string `json:"id"`
	SenderId   string `json:"sender_id"`
	ReceiverId string `json:"receiver_id"`
	Body       string `json:"body"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type ListMessage struct {
	Messages []*Message `json:"messages"`
	Count    int        `json:"count"`
}
type ListMessageReq struct {
	SenderId   string
	ReceiverId string
	FromDate   string
	ToDate     string
}
