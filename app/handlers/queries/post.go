package queries

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-pet-crud/ent"
	"go-pet-crud/ent/post"
)

type PostQueryHandler struct {
	client *ent.Client
}

func NewPostQueryHandler(client *ent.Client) *PostQueryHandler {
	return &PostQueryHandler{
		client: client,
	}
}

func (h *PostQueryHandler) Index(c *gin.Context) {
	p, err := h.client.Post.Query().All(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"posts": p})
}

func (h *PostQueryHandler) Show(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	p, err := h.client.Post.Query().
		Where(post.ID(id)).
		Only(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"post": p})
}
