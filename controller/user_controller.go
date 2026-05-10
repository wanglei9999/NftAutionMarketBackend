package controller

import (
	"NftAutionMarketBackend/database"
	"crypto/rand"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/gin-gonic/gin"
	"github.com/waglei9999/NftAutionMarketBackend/auth"
	"github.com/waglei9999/NftAutionMarketBackend/model"
	"github.com/waglei9999/NftAutionMarketBackend/database"
)

func NewUserController() *UserController {
	return &UserController{
		db:database.Db
	}
}	

type UserController struct {
	db *gorm.DB
}