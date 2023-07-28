package controllers

import (
	"net/http"
	"time"

	"github.com/BEERBAL-HASH/Vegapay-Limit-Offer/config"
	"github.com/BEERBAL-HASH/Vegapay-Limit-Offer/models"
	"github.com/gin-gonic/gin"
)

type CreateAccountInput struct {
	CustomerID uint `json:"customer_id" binding:"required"`
}

type CreateLimitOfferInput struct {
	AccountID           uint      `json:"account_id" binding:"required"`
	LimitType           string    `json:"limit_type" binding:"required"`
	NewLimit            float32   `json:"new_limit" binding:"required"`
	OfferActivationTime time.Time `json:"offer_activation_time" binding:"required"`
	OfferExpiryTime     time.Time `json:"offer_expiry_time" binding:"required"`
}

var LimitTypeChecker = []string{"ACCOUNT_LIMIT", "PER_TRANSACTION_LIMIT"}

func CreateLimitOffer(ctx *gin.Context) {
	var input CreateLimitOfferInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	flag := false
	for i := 0; i < len(LimitTypeChecker); i++ {
		if input.LimitType == LimitTypeChecker[i] {
			flag = true
		}
	}
	if !flag {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "LimitType should be ACCOUNT_LIMIT / PER_TRANSACTION_LIMIT"})
		return
	}
	if input.OfferActivationTime.Before(time.Now()) || input.OfferExpiryTime.Before(input.OfferActivationTime) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Offer Activation time or Expiry time isn't correct"})
		return
	}
	var validate models.Account

	err := config.DBconn.Where("account_id = ?", input.AccountID).First(&validate).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Account id doesn't exist"})
		return
	}

	limitOffer := models.LimitOffer{
		AccountID:           input.AccountID,
		NewLimit:            input.NewLimit,
		OfferActivationTime: input.OfferActivationTime,
		OfferExpiryTime:     input.OfferExpiryTime,
		LimitType:           input.LimitType,
		Status:              "PENDING",
	}

	savedOffer, err := limitOffer.SaveLimitOffer()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": savedOffer,
		"message": "Offer Created"})
}

func CreateAccount(ctx *gin.Context) {
	var input CreateAccountInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account := models.Account{
		CustomerID: input.CustomerID,
	}
	savedAccount, err := account.SaveAccount()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": savedAccount.AccountID,
		"message": "Account Created"})
}
