package dao

import (
	"context"
	"record-back/models"
	"time"
)

func SaveRecord(data *models.TradeStore) error {
	opCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := RecordCol.InsertOne(opCtx, &data)
	if err != nil {
		return err
	}

	return nil
}
