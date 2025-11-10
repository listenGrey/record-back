package utils

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math/big"
	"time"
)

func SToD128(s string) primitive.Decimal128 {
	val, _ := primitive.ParseDecimal128(s)
	return val
}

func D128ToS(d primitive.Decimal128) string {
	return d.String()
}

func SToTime(str string) time.Time {
	t, _ := time.Parse(time.RFC3339, str)

	return t
}

func TimeToS(t time.Time) string {
	return t.Format(time.RFC3339)
}

func D128Sub(price1, price2 primitive.Decimal128) primitive.Decimal128 {
	p1, _ := new(big.Float).SetString(price1.String())
	p2, _ := new(big.Float).SetString(price2.String())

	diff := new(big.Float).Sub(p2, p1)
	return SToD128(diff.String())
}
