package model

type Account struct {
	Orm                 `json:"orm"`
	Name                string
	MailConfig          []MailServer `json:"mail_config"`
	DefaultMailConfigID uint         `json:"default_mail_config_id"`
}

type MailServer struct {
	Orm       `json:"orm"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	From      string `json:"from"`
	Pass      string `json:"pass"`
	Server    string `json:"server"`
	Password  string `json:"password"`
	AccountID uint   `json:"account_id"`
}
