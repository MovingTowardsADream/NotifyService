package entity

type Channel struct {
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type Channels struct {
	EmailChannel Channel `json:"email"`
	PhoneChannel Channel `json:"phone"`
}

type RequestNotification struct {
	UserID     string   `json:"userId"`
	NotifyType string   `json:"notifyType"`
	Channels   Channels `json:"channels"`
}
