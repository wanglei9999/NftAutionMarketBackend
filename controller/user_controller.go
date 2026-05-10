package controller

import (
	"NftAutionMarketBackend/database"
	"crypto/rand"
	"fmt"
	"net/http"
	"strings"
	"time"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common"

	"github.com/gin-gonic/gin"
	"NftAutionMarketBackend/auth"
	"NftAutionMarketBackend/model"
	"NftAutionMarketBackend/database"
)

func NewUserController() *UserController {
	return &UserController{
		db:database.DB
	}
}	

type UserController struct {
	db *gorm.DB
}

type NonceRequest struct {
	Address string `json:"address" binding:"required"`
}

type LoginRequest struct {
	Address string `json:"address" binding:"required"`
	Nonce string `json:"nonce" binding:"required"`
	Signature string `json:"signature" binding:"required"`
}

func randomNonce() (string, error) {
	nonceBytes := make([]byte, 16)
	_, err := rand.Read(nonceBytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(nonceBytes), nil
}

func (uc *UserController) GetNonceHandler(c *gin.Context) {
	var req NonceRequest
	if err := c.ShouldBindJSON(&req); err != nil {{
		c.JSON(http.StatusBadRequest,gin.H{"error": "Invalid request"})
	}

	address := strings.ToLower(req.Address)
	if !common.IsHexAddress(address) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Ethereum address"})
		return
	}
	nonce, err := randomNonce()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate nonce"})
		return
	}
	auth.SetNonce(address, nonce, time.Now().Add(5*time.Minute))
	c.JSON(http.StatusOK, gin.H{"nonce": nonce})
}

func (uc *UserController) LoginHandler(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	address := strings.ToLower(req.Address)
	if !common.IsHexAddress(address) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Ethereum address"})
		return
	}
	storedNonce, exists := auth.GetNonce(address)
	if !exists || storedNonce != req.Nonce {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid nonce"})
		return
	}
	valid, err := auth.VerifySignature(address, req.Nonce, req.Signature)
	if err != nil || !valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid signature"})
		return
	}
	auth.DeleteNonce(address)
	// response token
	token, err := auth.GenerateToken(address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	var user model.User
	result := uc.db.Where("address = ?", address).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			user = model.User{Address: address}
			if err := uc.db.Create(&user).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
				return
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user, "token": token})
}