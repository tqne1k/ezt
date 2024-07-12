package bootstrap

import (
	"eztrust/domain"
	"eztrust/infra/db"
	"log"

	"gorm.io/gorm"
)

func Connect(env *Env) *gorm.DB {
	db, err := db.NewPgConnection(env.DbHost, env.DbPort, env.DbUser, env.DbPasswd, env.DbName)

	if err != nil {
		defer log.Println("Connection to Database Failed")
		log.Fatal(err.Error())
	}

	return db
}

func InitDabaBase(env *Env) error {
	err := db.InitDatabase(env.DbHost, env.DbPort, env.DbUser, env.DbPasswd, env.DbName)
	if err != nil {
		defer log.Println("Can not init database")
		log.Fatal(err.Error())
	}
	return nil
}

func Migrate(db *gorm.DB) {
	// db migration jobs
	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.Department{})
	db.AutoMigrate(&domain.Device{})
	db.AutoMigrate(&domain.Network{})
	db.AutoMigrate(&domain.DepartmentNetwork{})
}
