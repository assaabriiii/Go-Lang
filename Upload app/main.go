package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

const uploadDir = "/tmp/uploads"

func main() {
	err := os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		fmt.Println("Failed to create upload directory: ", err)
		return
	}

	router := gin.Default()

	router.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Error getting the file"})
			return
		}

		filePath := filepath.Join(uploadDir, file.Filename)
		err = c.SaveUploadedFile(file, filePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to list uploaded files"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"Message": "File uploaded successfully",
			"File":    file.Filename,
		})
	})

	router.GET("/files", func(c *gin.Context) {
		files, err := os.ReadDir(uploadDir)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to list files"})
			return
		}

		var fileNames []string
		for _, file := range files {
			fileNames = append(fileNames, file.Name())
		}

		c.JSON(http.StatusOK, gin.H{
			"Files": fileNames,
		})
	})

	router.GET("/download/:filename", func(c *gin.Context) {
		fileName := c.Param("filename")
		filePath := filepath.Join(uploadDir, fileName)

		_, err := os.Stat(filePath)
		if os.IsNotExist(err) {
			c.JSON(http.StatusNotFound, gin.H{"Error": "File not found"})
			return
		}

		c.File(filePath)
	})

	router.Run(":8080")
}
