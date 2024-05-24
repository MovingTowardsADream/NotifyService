package entity

type Preference struct {
	NotifyType string `json:"notifyType"`
	Channel    string `json:"channel"`
	Approval   bool   `json:"approval"`
}

type UserPreferences struct {
	UserID      string       `json:"userId"`
	Preferences []Preference `json:"preferences"`
}
