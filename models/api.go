package models

type TradeReq struct {
	OpenTime   string `json:"open_time"`
	CloseTime  string `json:"close_time"`
	Symbol     string `json:"symbol"`
	Side       string `json:"side"` // long / short
	Leverage   int    `json:"leverage"`
	OpenPrice  string `json:"open_price"`
	ClosePrice string `json:"close_price"`
	Fee        string `json:"fee"`
	Funding    string `json:"funding"`
	PnlGross   string `json:"pnl_gross"`
}
