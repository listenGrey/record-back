package logic

import (
	"fmt"
	"record-back/models"
)

func SaveTrade(form models.TradeReq) error {
	fmt.Println(form)
	return nil
}
