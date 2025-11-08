package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type TradeStore struct {
	OpenTime   time.Time            `bson:"open_time"`
	CloseTime  time.Time            `bson:"close_time"`
	Symbol     string               `bson:"symbol"`
	Side       string               `bson:"side"` // long / short
	Leverage   int                  `bson:"leverage"`
	OpenPrice  primitive.Decimal128 `bson:"open_price"`
	ClosePrice primitive.Decimal128 `bson:"close_price"`
	Fee        primitive.Decimal128 `bson:"fee"`
	Funding    primitive.Decimal128 `bson:"funding"`
	PnlGross   primitive.Decimal128 `bson:"pnl_gross"`
}
