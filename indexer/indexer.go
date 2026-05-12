package indexer


import (
	"context"
	"errors"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"


	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"NftAutionMarketBackend/model"
	"NftAutionMarketBackend/database"
	"gorm.io/gorm"

)







//获取连接客户端
func getConnect() *ethclient.Client {
	// 从环境变量读取，测试/主网一键切换
	rpcURL := os.Getenv("ETH_RPC_URL")

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatal(err)
	}
	return client
}