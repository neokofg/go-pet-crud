package commands

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-pet-crud/ent"
)

type PostHandler struct {
	client *ent.Client
}

func NewPostHandler(client *ent.Client) *PostHandler {
	return &PostHandler{
		client: client,
	}
}

type CreateRequest struct {
	Title  string `json:"title" binding:"required"`
	Body   string `json:"body" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateRequest struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author string `json:"author"`
}

func (h *PostHandler) Create(c *gin.Context) {
	var req CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	_, err := h.client.Post.Create().
		SetBody(req.Body).
		SetAuthor(req.Author).
		SetTitle(req.Title).
		Save(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "post created"})
}

func (h *PostHandler) Update(c *gin.Context) {
	var req UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	var builder = h.client.Post.UpdateOneID(id)
	if req.Title != "" {
		builder.SetTitle(req.Title)
	}
	if req.Body != "" {
		builder.SetBody(req.Body)
	}
	if req.Author != "" {
		builder.SetAuthor(req.Author)
	}
	_, err = builder.Save(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "post updated"})
}

func (h *PostHandler) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if err := h.client.Post.DeleteOneID(id).Exec(c.Request.Context()); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "post deleted"})
}
