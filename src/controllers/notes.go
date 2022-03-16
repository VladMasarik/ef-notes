package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/VladMasarik/ef-notes/src/database"
	"github.com/VladMasarik/ef-notes/src/models"
)

type NotesAPI struct{}

// func handleServerError(ctx *gin.Context, message string) {

// }

func (nt NotesAPI) List(ctx *gin.Context) {
	var notes []models.Note

	if err := database.DB.Find(&notes).Error; err != nil {
		ctx.Error(err)
		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "Failed to get data from DB",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"data": notes,
	})
}

func (nt NotesAPI) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	var note models.Note

	if err := database.DB.First(&note, id).Error; err != nil {
		ctx.Error(err)
		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "Failed to get data from DB",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"data": note,
	})
}

func (nt NotesAPI) Post(ctx *gin.Context) {
	var note models.Note
	ctx.BindJSON(&note)

	if err := database.DB.Create(&note).Error; err != nil {
		ctx.Error(err)
		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "Failed to create data in DB",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"data": note,
	})
}

func (nt NotesAPI) Put(ctx *gin.Context) {
	var note models.Note
	ctx.BindJSON(&note)

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.Error(err)
		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "Failed to parse ID",
		})
		return
	}
	note.ID = uint(id)

	if err := database.DB.Save(&note).Error; err != nil {
		ctx.Error(err)
		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "Failed to update data in DB",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"data": note,
	})
}

func (nt NotesAPI) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := database.DB.Delete(&models.Note{}, id).Error; err != nil {
		ctx.Error(err)
		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "Failed to delete data in DB",
		})
		return
	}
	ctx.Status(204) // 204 is for Success with no content in the response
}
