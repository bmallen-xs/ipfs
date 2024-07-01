package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func startAPI(db *sqlx.DB) {
	router := gin.Default()

	router.GET("/metadata", func(c *gin.Context) {
		metadata, err := getAllMetadata(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, metadata)
	})

	router.GET("/metadata/:cid", func(c *gin.Context) {
		cid := c.Param("cid")
		metadata, err := getMetadataByCID(db, cid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, metadata)
	})

	router.Run(":8080")
}
