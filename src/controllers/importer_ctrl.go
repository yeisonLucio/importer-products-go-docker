package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/importer/src/repositories"
	"github.com/importer/src/usecases"
)

func CreateProducts(ctx *gin.Context) {

	file, _ := ctx.FormFile("products")
	extension := filepath.Ext(file.Filename)
	fileName := uuid.New().String() + extension
	filePath := fmt.Sprintf("./storage/%v.csv", fileName)

	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	if err := repositories.NewSummary(filePath); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	usecases.ProcessFile(filePath)

	ctx.JSON(http.StatusOK, gin.H{
		"res": true,
	})
}
