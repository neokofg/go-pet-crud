package router

import (
	"github.com/gin-gonic/gin"
	"go-pet-crud/app/handlers/commands"
	"go-pet-crud/app/handlers/queries"
	"go-pet-crud/ent"
)

func InitRouter(r *gin.Engine, client *ent.Client) {
	postHandler := commands.NewPostHandler(client)
	postQueryHandler := queries.NewPostQueryHandler(client)

	posts := r.Group("/posts")
	posts.POST("/", postHandler.Create)
	posts.PUT("/:id", postHandler.Update)
	posts.DELETE("/:id", postHandler.Delete)

	posts.GET("/", postQueryHandler.Index)
	posts.GET("/:id", postQueryHandler.Show)
}
