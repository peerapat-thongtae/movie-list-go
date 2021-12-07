package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func goDotEnvVariable(key string) string {

  // load .env file
  err := godotenv.Load(".env")

  if err != nil {
    log.Fatalf("Error loading .env file")
  }

  return os.Getenv(key)
}


func (h *TodoHandler) GetAllTodo(c *gin.Context) {
	todos := []Todo{}

	h.DB.Table("todo").Limit(10).Find(&todos)

	c.JSON(http.StatusOK, todos)
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	h := TodoHandler{}
	h.Initialize()

	r.GET("/todos", h.GetAllTodo)

	return r
}

type Todo struct {
	Id        string `gorm:"primary_key" json:"id"`
	MediaId uint `gorm:"column:mediaId" json:"tmdbId"`
	UserId  uint `gorm:"column:userId" json:"userId"`
	Status       string    `json:"status"`
}


type TodoHandler struct {
	DB *gorm.DB
}

func (h *TodoHandler) Initialize() {
	conn := goDotEnvVariable("DATABASE_URI")
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	h.DB = db
}

func main() {
	r := setupRouter()
	
	r.Run()
}