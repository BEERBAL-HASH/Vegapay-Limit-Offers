package controllers

import (
	"net/http"
	"strconv"

	"github.com/BEERBAL-HASH/Vegapay-Limit-Offer/config"
	"github.com/BEERBAL-HASH/Vegapay-Limit-Offer/models"
	"github.com/gin-gonic/gin"
)

func GetAccountById(ctx *gin.Context) {
	db := config.DBconn
	var account models.Account

	paramId := ctx.Param("account_id")
	id, _ := strconv.Atoi(paramId)

	err := db.Where("account_id=?", id).First(&account).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"data": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": account})
}

func GetLimitOffers(ctx *gin.Context) {
	db := config.DBconn
	var limitOffers []models.LimitOffer

	paramId := ctx.Param("account_id")
	id, _ := strconv.Atoi(paramId)

	db.Where("account_id = ? AND status=?", id, "PENDING").Find(&limitOffers)
	if len(limitOffers) < 1 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "data not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": &limitOffers})
}
