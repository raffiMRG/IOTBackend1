package Connection

import (
	// "gorm.io/driver/postgres"
	"backendIOT1/Model/AbsenMhsModel"
	"backendIOT1/Model/ListMhsModel"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() (*gorm.DB, error) {
	// MYSQL DSN
	// dsn := "root:@tcp(localhost:3306)/db_buku_tamu?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "mrg:123123123@tcp(localhost:3306)/cobarelasi2?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// POSGRESQL DSN
	// dsn := "host=aws-0-ap-southeast-1.pooler.supabase.com user=postgres.wbxuqhuitgesreybeikb password=Sasdswsz1234 dbname=postgres port=6543 sslmode=require TimeZone=Asia/Jakarta"
	// dsn := "host=aws-0-ap-southeast-1.pooler.supabase.com port=6543 user=postgres.wbxuqhuitgesreybeikb password=Sasdswsz1234 dbname=postgres sslmode=require"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	DB = db
	db.AutoMigrate(&ListMhsModel.ListMhs{}, &AbsenMhsModel.ListAbsen{})
	return db, nil
}
