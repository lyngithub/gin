package forms

type UserStatus struct {
	CallId    string `mapstructure:"callId" json:"callId"`       // callid
	Security  string `mapstructure:"security" json:"security"`   // 签名
	Timestamp int    `mapstructure:"timestamp" json:"timestamp"` // 时间戳
	User      string `mapstructure:"user" json:"user"`           // 用户名
	Status    string `mapstructure:"user" json:"status"`         // 用户状态
	Ip        string `mapstructure:"ip" json:"ip"`
	Version   string `mapstructure:"version" json:"version"`
}

type Exception struct {
	CallId string `mapstructure:"callId" json:"callId"` // callid
}
