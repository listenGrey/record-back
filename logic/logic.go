package logic

import (
	"record-back/dao"
	"record-back/models"
	"record-back/utils"
)

func SaveTrade(form models.TradeReq) error {
	store := &models.TradeStore{
		OpenTime:   utils.SToTime(form.OpenTime),
		CloseTime:  utils.SToTime(form.CloseTime),
		Symbol:     form.Symbol,
		Side:       form.Side,
		Leverage:   form.Leverage,
		OpenPrice:  utils.SToD128(form.OpenPrice),
		ClosePrice: utils.SToD128(form.ClosePrice),
		Fee:        utils.SToD128(form.Fee),
		Funding:    utils.SToD128(form.Funding),
		PnlGross:   utils.SToD128(form.PnlGross),
	}

	err := dao.SaveRecord(store)
	if err != nil {
		return err
	}

	return nil
}
