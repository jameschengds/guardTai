package api

//PushTX ...
type PushTX struct {
	From string `form:"from"`
	To   string `form:"to"`
	Pk   string `form:"pk"`
	Memo string `form:"memo"`
}

type NewAccountReq struct {
	AccountName string `json:"account_name"`
	PubKey      string `json:"pub_key"`
}
