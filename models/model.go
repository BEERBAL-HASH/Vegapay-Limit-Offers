package models

import "github.com/BEERBAL-HASH/Vegapay-Limit-Offer/config"

func (entry *Account) SaveAccount() (*Account, error) {
	err := config.DBconn.Create(&entry).Error
	if err != nil {
		return &Account{}, err
	}
	return entry, nil
}

func (entry *LimitOffer) SaveLimitOffer() (*LimitOffer, error) {
	err := config.DBconn.Create(&entry).Error
	if err != nil {
		return &LimitOffer{}, err
	}
	return entry, nil
}
