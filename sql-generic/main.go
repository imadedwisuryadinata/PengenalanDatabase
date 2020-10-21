package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/FadhlanHawali/Digitalent-Kominfo_Introduction-Database-1/sql-generic/config"
	"github.com/FadhlanHawali/Digitalent-Kominfo_Introduction-Database-1/sql-generic/database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func main() {
	cfg, err := getConfig()
	if err != nil {
		log.Println(err)
		return
	}

	db, err := connect(cfg.Database)
	if err != nil {
		log.Println(err)
		return
	}

	// database.InsertCustomer(database.Customer{
	// 	FirstName:    "budi",
	// 	LastName:     "arjo",
	// 	NpwpId:       "npwp123",
	// 	Age:          20,
	// 	CustomerType: "premium",
	// 	Street:       "bandung",
	// 	City:         "jakarta",
	// 	State:        "Jawa Barat",
	// 	Zip_code:     "1233",
	// 	PhoneNumber:  "08111113214",
	// }, db)

	//database.GetCustomers(db)
	database.UpdateCustomer(45, 1, db)
	//database.DeleteCustomer(1, db)
}

func getConfig() (config.Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")
	viper.SetConfigName("config.yml")

	if err := viper.ReadInConfig(); err != nil {
		return config.Config{}, err
	}

	var cfg config.Config
	err := viper.Unmarshal(&cfg)
	if err != nil {
		return config.Config{}, err
	}

	return cfg, nil
}

func connect(cfg config.Database) (*sql.DB, error) {
	db, err := sql.Open(cfg.Driver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DbName, cfg.Config))
	if err != nil {
		return nil, err
	}

	log.Println("db successfully connected")
	return db, nil
}
