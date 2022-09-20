package config

import (
	"backend/utils"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {

	dbConfig := DBConfig{
		Host:    "localhost",
		Port:     3306,
		User:     "root",
		Password: "secret",
		DBName:   "diary",
	}
	return &dbConfig
}
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
var err error
func ConnectToDB(){
	config, err := utils.LoadConfig("./../utils");
	fmt.Println(config);
	if err !=nil {
		utils.SugarLogger.Error(err);
		utils.SugarLogger.Error("error in loading config file with viper");
	}


	DB, err = gorm.Open(config.Db.DbDriver, DbURL(BuildDBConfig()))
	if err != nil {
		utils.SugarLogger.Error("DataBase is not connected",zap.Error(err))
	}
}
