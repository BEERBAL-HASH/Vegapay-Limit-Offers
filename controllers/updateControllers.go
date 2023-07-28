package controllers

import (
	"net/http"
	"time"

	"github.com/BEERBAL-HASH/Vegapay-Limit-Offer/config"
	"github.com/BEERBAL-HASH/Vegapay-Limit-Offer/models"
	"github.com/gin-gonic/gin"
)

type UpdateLimitOfferStatusInput struct {
	LimitOfferID uint   `json:"Limit_offer_id" binding:"required"`
	Status       string `json:"status" binding:"required"`
}

var StatusChecker = []string{"ACCEPTED", "REJECTED"}

func UpdateLimitOfferStatus(ctx *gin.Context) {

	var input UpdateLimitOfferStatusInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	flag := false
	for i := 0; i < len(StatusChecker); i++ {
		if input.Status == StatusChecker[i] {
			flag = true
		}
	}
	if !flag {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Status isn't passed in a correct way"})
		return
	}

	db := config.DBconn
	tx := db.Begin()

	var limitOffer models.LimitOffer
	var account models.Account
	err := tx.Where("limit_offer_id=? AND status=?", input.LimitOfferID, "PENDING").First(&limitOffer).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error(), "error message": "offer doesn't exist any more"})
		return
	}

	err = tx.Where("account_id=?", limitOffer.AccountID).First(&account).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if limitOffer.OfferExpiryTime.Before(time.Now()) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "offer has been expired"})
		return
	}

	if input.Status == "ACCEPTED" {

		if limitOffer.LimitType == "ACCOUNT_LIMIT" {
			if account.AccountLimit > limitOffer.NewLimit {
				ctx.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error(), "message": "OOPS.....Not Eligible"})
				return
			}
			account.LastAccountLimit = account.AccountLimit
			account.AccountLimitUpdateTime = time.Now()
			account.AccountLimit = limitOffer.NewLimit
		} else if limitOffer.LimitType == "per_transaction_limit" {
			if account.PertransactionLimit > limitOffer.NewLimit {
				ctx.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error(), "message": "OOPS....Not Eligible"})
				return
			}
			account.LastPertransactionLimit = account.PertransactionLimit
			account.PertransactionLimitUpdateTime = time.Now()
			account.PertransactionLimit = limitOffer.NewLimit
		}

		err = tx.Save(&account).Error
		if err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	limitOffer.Status = input.Status
	err = tx.Save(&limitOffer).Error
	if err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tx.Commit()

	ctx.JSON(http.StatusOK, gin.H{"data": limitOffer, "message": "Response Recorded"})
}
