package handler

import (
	"github.com/gin-gonic/gin"
	"myskill-api/datastore"
	"net/http"
)

type User interface {
	GetUserByID(gctx *gin.Context)
}

type UserHandler struct {
	userStore datastore.User
}

func NewUserHandler(db datastore.User) *UserHandler {
	return &UserHandler{
		userStore: db,
	}
}

func (u *UserHandler) GetUserByID(gctx *gin.Context) {
	gctx.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}
