package models

// Settings app settings
type Settings struct {
	AppParams  Params       `json:"app"`
	UserParams UserSettings `json:"user_params"`
}

// Params contains params of server meta data
type Params struct {
	ServerName string `json:"server_name"`
	PortRun    string `json:"port_run"`
	LogFile    string `json:"log_file"`
	ServerURL  string `json:"server_url"`
}

type UserSettings struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
