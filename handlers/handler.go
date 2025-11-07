package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"record-back/logic"
	"record-back/models"
)

func CreateTrade(c *gin.Context) {
	var trade models.TradeReq
	if err := c.ShouldBindJSON(&trade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := logic.SaveTrade(trade)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库插入失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "保存成功"})
}
