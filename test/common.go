package test

import (
	"github.com/joho/godotenv"
	"github.com/triaton/forum-backend-echo/database"
	"github.com/triaton/forum-backend-echo/migrations"
	"github.com/triaton/forum-backend-echo/models"
	"log"
	"os"
)

func InitTest() {
	err := godotenv.Load(os.ExpandEnv("$GOPATH/src/github.com/triaton/forum-backend-echo/test.env"))
	if err != nil {
		log.Fatal("failed to load test env config: ", err)
	}
	db := database.GetInstance()
	// TODO: remove all tables before test
	db.DropTable("migrations")
	db.DropTableIfExists(&models.User{})
	db.DropTableIfExists(&models.Blog{})
	db.DropTableIfExists(&models.Comment{})
	m := migrations.GetMigrations(db)
	err = m.Migrate()
	if err != nil {
		log.Fatal("failed to run db migration: ", err)
	}
}
