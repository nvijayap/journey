package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	uhd, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	err = godotenv.Load(uhd + "/.env")
	if err != nil {
		log.Fatal(err)
	}

	pghost, pgport := os.Getenv("pghost"), os.Getenv("pgport")
	pgun, pgpw := os.Getenv("pgun"), os.Getenv("pgpw")
	pgdb := os.Getenv("pgdb")

	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	r.GET("/connect", func(c *gin.Context) {
		connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			pgun, pgpw, pghost, pgport, pgdb)
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		}
		if err = db.Ping(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "connected",
		})
	})

	r.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })

	r.Run(":8989")
}
