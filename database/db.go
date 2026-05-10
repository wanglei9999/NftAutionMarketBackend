package database


import (
	"log"
	"time"
	"github.com/waglei9999/NftAutionMarketBackend/model"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)


var DB *gorm.DB

func InitializeDatabase() {
	//connect to database
	dsn := "root:password@tcp(localhost:3306)/nft_auction_market?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	//最大连接数
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to get database instance:", err)
	}
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

	//自动建表
	err = DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	err = DB.AutoMigrate(&model.Auction{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	err = DB.AutoMigrate(&model.Bid{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	err = DB.AutoMigrate(&model.NFT{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	err = DB.AutoMigrate(&model.SyncStatus{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	log.Println("Database connection established and migrated successfully.")
}

