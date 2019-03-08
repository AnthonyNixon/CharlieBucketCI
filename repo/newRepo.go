package repo

import (
	"github.com/AnthonyNixon/CharlieBucketCI/database"
	"github.com/AnthonyNixon/CharlieBucketCI/types"
	"github.com/gin-gonic/gin"
)

func NewRepo(c *gin.Context) {
	var repo types.Repo
	c.BindJSON(&repo)

	database, err := database.GetConnection()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer database.Close()

	stmt, err := database.Prepare("insert into repos (org, repo, loadbalancer, domain) values (?,?,?,?);")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	_, err = stmt.Exec(repo.Org, repo.Name, repo.Loadbalancer, repo.Domain)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"status": "created"})
}
