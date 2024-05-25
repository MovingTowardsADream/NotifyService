package entity

type NotificationRequest struct {
	UserID     string             `json:"userId"`
	NotifyType string             `json:"notifyType"`
	Channels   map[string]Channel `json:"channels"`
}

type Channel struct {
	Subject string `json:"subject"`
	Body    string `json:"body"`
}
