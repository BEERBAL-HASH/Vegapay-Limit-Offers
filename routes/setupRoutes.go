package routes

import (
	"github.com/BEERBAL-HASH/Vegapay-Limit-Offer/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	account := router.Group("/account")
	offer := router.Group("/offer")

	//handling update request on path /offer/update, takes limit_offer_id in req body
	offer.PUT("/update", controllers.UpdateLimitOfferStatus)

	//handling post request on path /offer/create
	offer.POST("/create", controllers.CreateLimitOffer)

	//handling get request on path /offer/view/1
	offer.GET("/view/:account_id", controllers.GetLimitOffers)

	//handling post request on path /account/create
	account.POST("/create", controllers.CreateAccount)

	//handling get request on path /account/view
	account.GET("/view/:account_id", controllers.GetAccountById)

}
