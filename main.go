package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	utils "github.com/masharpik/bdProject/utils"
)

type dbEnv struct {
	userTitle     string
	passwordTitle string
	netlocTitle   string
	portTitle     string
	dbnameTitle   string

	user     string
	password string
	netloc   string
	port     string
	dbname   string
}

func getEnvDB(db dbEnv) dbEnv {
	var exists bool

	db.user, exists = os.LookupEnv(db.userTitle)
	if !exists {
		log.Fatalf("Не нашлось %s в .env", db.userTitle)
	}

	db.password, exists = os.LookupEnv(db.passwordTitle)
	if !exists {
		log.Fatalf("Не нашлось %s в .env", db.passwordTitle)
	}

	db.netloc, exists = os.LookupEnv(db.netlocTitle)
	if !exists {
		log.Fatalf("Не нашлось %s в .env", db.passwordTitle)
	}

	db.port, exists = os.LookupEnv(db.portTitle)
	if !exists {
		log.Fatalf("Не нашлось %s в .env", db.portTitle)
	}

	db.dbname, exists = os.LookupEnv(db.dbnameTitle)
	if !exists {
		log.Fatalf("Не нашлось %s в .env", db.dbnameTitle)
	}

	return db
}

func getDialector(db dbEnv) string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		db.user, db.password, db.netloc, db.port, db.dbname)
}

func getDbConnection() *gorm.DB {
	dbOpts := dbEnv{
		userTitle:     "DB_USER",
		passwordTitle: "DB_PASS",
		netlocTitle:   "DB_HOST",
		portTitle:     "DB_PORT",
		dbnameTitle:   "DB_NAME",
	}

	dbOpts = getEnvDB(dbOpts)

	dialector := getDialector(dbOpts)

	// Установка соединения с базой данных
	db, err := gorm.Open(postgres.Open(dialector), &gorm.Config{})
	if err != nil {
		log.Fatalln("Не удалось установить соединение с БД: ", err)
	}

	return db
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Нет .env файла: ", err)
	}

	db := getDbConnection()
	sqlDb, err := db.DB()
	if err != nil {
		log.Fatalln("Не удалось получить объект sqlDB: ", err)
	}
	defer sqlDb.Close()

	if err = utils.FillDB(db); err != nil {
		log.Fatalln("Не удалось заполнить БД: ", err)
	}
}
