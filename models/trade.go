package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Trade struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	OpenTime   time.Time          `bson:"openTime" json:"open_time"`
	CloseTime  time.Time          `bson:"closeTime" json:"close_time"`
	Symbol     string             `bson:"symbol" json:"symbol"`
	Side       string             `bson:"side" json:"side"` // long / short
	Leverage   int                `bson:"leverage" json:"leverage"`
	OpenPrice  float64            `bson:"openPrice" json:"open_price"`
	ClosePrice float64            `bson:"closePrice" json:"close_price"`
	Fees       []float64          `bson:"fees" json:"fees"`
	Funding    float64            `bson:"funding" json:"funding"`
	PnlGross   float64            `bson:"pnlGross" json:"pnl_gross"`
	CreatedAt  time.Time          `bson:"createdAt" json:"created_at"`
}
