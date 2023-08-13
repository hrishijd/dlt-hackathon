package db

import (
	"fmt"
	"os"
	"time"

	"github.com/m/v2/config"
	"github.com/m/v2/db/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDataBase struct {
	DB *gorm.DB
}

var PostgresDBGlobalInstance *PostgresDataBase

func InitializeDB() {
	dBConfig := config.ReadDBConfig()
	dbURL := "host=" + dBConfig.Host + " user=" + dBConfig.User + " password=" + dBConfig.Password + " dbname=" + dBConfig.Name + " port=" + dBConfig.Port + " sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error unable to open DB: %s", err)
		os.Exit(1)
	}

	if err := db.AutoMigrate(&models.Supply{}); err != nil {
		fmt.Printf("Failed to automigrate tables %s", err)
		os.Exit(1)
	}

	PostgresDBGlobalInstance = &PostgresDataBase{
		DB: db,
	}
}

func (db *PostgresDataBase) CreateSupply(supply *models.Supply) error {

	if err := PostgresDBGlobalInstance.DB.Create(&supply); err.Error != nil {
		return err.Error
	}

	return nil
}

func (db *PostgresDataBase) UpdateSupply(supply *models.Supply) error {

	result := PostgresDBGlobalInstance.DB.
		Model(models.Supply{}).
		Where("supplier = ?", supply.Supplier).
		Updates(supply)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (db *PostgresDataBase) GetSupplyBySupplier(supplier string) (*models.Supply, error) {
	supply := models.Supply{}

	result := PostgresDBGlobalInstance.DB.
		Model(models.Supply{}).
		Where("supplier = ?", supplier).
		Limit(1).
		Take(&supply)
	if result.Error != nil {
		return &supply, result.Error
	}

	return &supply, nil
}

func (db *PostgresDataBase) GetSupplysSortedByPriority() ([]*models.Supply, error) {
	var PriorityList []*models.Supply

	result := db.DB.
		Model(&models.Supply{}).
		Where(
			"split_time + last_supply_time < ?",
			uint64(time.Now().Unix()),
		).
		Order(struct {
			LastSupplyTime uint64
		}{}).
		Find(&PriorityList)
	if result.Error != nil {
		return nil, fmt.Errorf("CreatePriorityList: unable to get list from db %w", result.Error)
	}

	return PriorityList, nil
}
