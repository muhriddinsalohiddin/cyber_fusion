package models  
type Notification struct{
	Id string            `json:"id"`
	UserId string         `json:"user_id"`
	Type string             `json:"type"`
	Body string              `json:"body"`
	CreatedAt string         `json:"created_at"`
}
type List struct {
	Notifications []*Notification `json:"notifications"`
	Count     int        `json:"count"`
}