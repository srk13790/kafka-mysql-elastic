package handlers

import (
	"kafka-new/crawler"
	"kafka-new/kafka"
	"os"

	"github.com/gin-gonic/gin"
)

func CrawlDataFromURL(c *gin.Context) {
	url := c.Query("url")

	if url == "" {
		c.JSON(400, gin.H{
			"error": "Url is required",
		})
	}

	data, err := crawler.CrawlData(url)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to crawl data",
		})
	}

	file, err := os.Create("crawledData.txt")
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to create file",
		})
	}

	_, err = file.Write(data)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to write file",
		})
	}

	// MySQL connection to write data
	

	// Redis Connection to write data
	kafka.KafkaExec(data)

	c.JSON(200, gin.H{"message": "Data written successfully"})

}