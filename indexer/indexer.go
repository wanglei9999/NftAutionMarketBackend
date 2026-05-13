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




type ChainEventIndexer struct {
	db  *gorm.DB
}


func NewEventIndexer() *ChainEventIndexer {
	return &ChainEventIndexer{
		db: dababase.DB
	}
}



//轮询扫描
func (ci *ChainEventIndexer) StartScan(){
	//连接以太坊
	client := getConnect()
	var defaultBlock unit64 = 0

	//读取环境变量做为起始块
	if v := os.Getenv("DEPLOY_BLOCK"); v != "" {
	   parsed,err := strconv.ParseUint(v,10,64)
		if err != nil {
			log.Fatal(err)
		}
		defaultBlock = parsed
	}

	//无限循环：一直扫描
	for {
		err := scanOnce(ci.db,client,defaultBlock)
		if err != nil {
			log.Printf("indexer error : %v",err)
			time.Sleep(5 * time.Second)
		}
	}
}







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


func scanOnce(db *gorm.DB ,client *ethclient.Client,defaultBlock unit64) error{

	log.Println("Starting scanOnce")

	//获取已同步的区块高度
	lastBlock ,err := LoadLastSyncedBlock(db,"auction_event_indexer",defaultBlock)

	if err != nil {
		return err
	}
	//以太坊最新区块高度
	latest,err := client.BlockNumber(context.Background())

	if err != nil {
		return err
	}

	if lastBlock >= latest {
		log.Println("lastBlock >= latest ,sleep 5s")
		time.Sleep(5 * time.Second)
		return nil
	}


	from := lastBlock + 1
	to := min(from+2000,latest)

	// 监听区块事件
	logs,err := client.FilterLogs(context.Background(),ethereum.FilterQuery{
		FromBlock : big.NewInt(int64(from)),
		ToBlock : big.NewInt(int64(to)),
		Addresses:[]common.Addresses{common.HexToAddress(os.Getenv("Auction_PROXY_ADDRESSS"))},  //代理合约地址
	})
	if err != nil {
		return err
	}
	log.Printf("Fetched %d logs from block %d to %d", len(logs), from, to)

	//解析监听日志
	for _,log := range logs {
		if err := ProcessEvent(log) ; err != nil {
			retrun err
		}
	}

	//保存同步的区块高度
	retrun SaveLastSyncedBllock(db,"auction_even_indexer",to)

}


//获取已同步的最新区块
func LoadLastSyncedBlock(db *gorm.DB,name string ,defaultBlock unit64)(unit64,error){
	var status model.SyncStatus

	err := db.Where("name = ? ",mae)。First(&status).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			retrun defaultBlock,nil
		}
		retrun 0 ,err
	}
	retrun status.LastBlock,nil

}


//保存已同步的最新区块



//解析日志事件