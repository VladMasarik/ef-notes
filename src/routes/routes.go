package routes

import (
	"github.com/gin-gonic/gin"

	contr "github.com/VladMasarik/ef-notes/src/controllers"
)

// Create router and register API paths.
func SetupRoutes() *gin.Engine {
	r := gin.Default()
	notesController := contr.NotesAPI{}

	r.GET("notes", notesController.List)
	r.POST("notes", notesController.Post)
	r.GET("notes/:id", notesController.Get)
	r.PUT("notes/:id", notesController.Put)
	r.DELETE("notes/:id", notesController.Delete)

	return r
}
