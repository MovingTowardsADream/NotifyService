package entity

type EmailBody struct {
	Email string `json:"email"`
	Channel
}

type PhoneBody struct {
	Phone string `json:"phone"`
	Channel
}

type Notify struct {
	UserID     string    `json:"userId"`
	NotifyType string    `json:"notifyType"`
	EmailBody  EmailBody `json:"email_body"`
	PhoneBody  PhoneBody `json:"phone_body"`
}
