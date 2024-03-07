package app

import (
	"github.com/danjelhysenaj-dev/bookstore_auth-api/src/domain/access_token"
	"github.com/danjelhysenaj-dev/bookstore_auth-api/src/http"
	"github.com/danjelhysenaj-dev/bookstore_auth-api/src/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atHandler := http.NewAccessTokenHandler(access_token.NewService(db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("oauth/access_token", atHandler.Create)
	router.Run(":8080")
}
