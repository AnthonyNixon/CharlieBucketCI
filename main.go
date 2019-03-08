package main

import (
	"github.com/AnthonyNixon/CharlieBucketCI/database"
	"github.com/AnthonyNixon/CharlieBucketCI/loadbalancer"
	"github.com/AnthonyNixon/CharlieBucketCI/repo"
	"github.com/gin-gonic/gin"
)

func init() {
	database.Initialize()
}

func main() {
	r := gin.Default()
	r.POST("/repo", func(c *gin.Context) {
		repo.NewRepo(c)
	})

	r.POST("/loadbalancer", func(c *gin.Context) {
		loadbalancer.NewLoadbalancer(c)
	})

	r.Run() // Listen and serve on 0.0.0.0:8080
}
