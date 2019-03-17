package api

type PushTX struct {
	From string `form:"from"`
	To   string `form:"to"`
	Pk   string `form:"pk"`
	Memo string `form:"memo"`
}
