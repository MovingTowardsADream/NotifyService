package entity

type Preferences struct {
	Email EmailPreferences `json:"email"`
	Phone PhonePreferences `json:"phone"`
}

type EmailPreferences struct {
	NotifyType string `json:"notifyType"`
	Approval   bool   `json:"approval"`
}

type PhonePreferences struct {
	NotifyType string `json:"notifyType"`
	Approval   bool   `json:"approval"`
}

type RequestPreferences struct {
	UserID      string      `json:"userId"`
	Preferences Preferences `json:"preferences"`
}
