package test

import (
	"github.com/joho/godotenv"
	BlogModels "github.com/triaton/forum-backend-echo/blogs/models"
	"github.com/triaton/forum-backend-echo/database"
	UserModels "github.com/triaton/forum-backend-echo/users/models"
	"log"
	"os"
)

func LoadTestEnv() error {
	err := godotenv.Load(os.ExpandEnv("$GOPATH/src/github.com/triaton/forum-backend-echo/test.env"))
	if err != nil {
		log.Fatal("failed to load test env config: ", err)
	}
	return err
}

func InitTest() {
	err := LoadTestEnv()
	db := database.GetInstance()
	db.DropTable("migrations")
	db.DropTableIfExists(&UserModels.User{})
	db.DropTableIfExists(&BlogModels.Blog{})
	m := database.GetMigrations(db)
	err = m.Migrate()
	if err != nil {
		log.Fatal("failed to run db migration: ", err)
	}
}
