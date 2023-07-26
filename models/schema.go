package models

import (
	"time"
)

// account table cointains account info and Account table has one to many relationship with LimitOffer Table
type Account struct {
	AccountID                     uint         `json:"account_id" gorm:"primaryKey:autoincrement"`
	CustomerID                    uint         `json:"customer_id"`
	AccountLimit                  float32      `json:"account_limit"`
	PertransactionLimit           float32      `json:"per_transaction_limit"`
	LastAccountLimit              float32      `json:"last_account_limit"`
	LastPertransactionLimit       float32      `json:"last_per_transaction_limit"`
	AccountLimitUpdateTime        time.Time    `json:"account_limit_update_time" gorm:"autoUpdateTime:nano"`
	PertransactionLimitUpdateTime time.Time    `json:"per_transaction_limit_update_time" gorm:"autoUpdateTime:nano"`
	LimitOffers                   []LimitOffer `json:"limit_offers" gorm:"foreignKey:AccountID"`
}

// LimitOffer table cointains info of all the available limit offers for all accounts
type LimitOffer struct {
	AccountID           uint      `json:"account_id"`
	LimitOfferID        uint      `json:"offer_id" gorm:"primaryKey:autoincrement"`
	LimitType           string    `json:"limit_type"`
	NewLimit            float32   `json:"new_limit"`
	OfferActivationTime time.Time `json:"offer_activation_time" gorm:"index"`
	OfferExpiryTime     time.Time `json:"offer_expiry_time"`
	Status              string    `json:"status"`
}
