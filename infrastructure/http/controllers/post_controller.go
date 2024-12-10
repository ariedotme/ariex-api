package controllers

import (
	"net/url"

	"github.com/ariedotme/ariex-backend/application/services"
	"github.com/ariedotme/ariex-backend/application/usecases"
	"github.com/ariedotme/ariex-backend/domain/entities"
	"github.com/gofiber/fiber/v2"
)

type PostController struct {
	CreatePostUseCase *usecases.CreatePostUseCase
	PostService       *services.PostService
}

// CreatePost godoc
// @Summary Create a new blog post
// @Description Adds a new blog post in markdown format
// @Tags Posts
// @Accept json
// @Produce json
// @Param request body map[string]string true "Post Data"
// @Success 201 {object} entities.Post
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /admin/posts [post]
// @security BearerAuth
func (c *PostController) CreatePost(ctx *fiber.Ctx) error {
	var payload struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid payload"})
	}

	post, err := c.CreatePostUseCase.Execute(payload.Title, payload.Content)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(post)
}

// GetPosts godoc
// @Summary Get all blog posts
// @Description Retrieves a list of all blog posts
// @Tags Posts
// @Produce json
// @Success 200 {array} entities.Post
// @Failure 500 {object} map[string]string
// @Router /posts [get]
func (c *PostController) GetPosts(ctx *fiber.Ctx) error {
	posts, err := c.PostService.GetPosts()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(posts)
}

// GetPostByID godoc
// @Summary Get a blog post by ID
// @Description Retrieves a single blog post by its ID
// @Tags Posts
// @Produce json
// @Param id path string true "Post ID"
// @Success 200 {object} entities.Post
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /posts/{id} [get]
func (c *PostController) GetPostByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	post, err := c.PostService.GetPostByID(id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Post not found"})
	}

	return ctx.Status(fiber.StatusOK).JSON(post)
}

// GetPostByNormalizedTitle godoc
// @Summary Get a blog post by normalized title
// @Description Retrieves a single blog post by its normalized title
// @Tags Posts
// @Produce json
// @Param normalized_title path string true "Post Normalized Title"
// @Success 200 {object} entities.Post
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /posts/title/{normalized_title} [get]
func (c *PostController) GetPostByNormalizedTitle(ctx *fiber.Ctx) error {
	normalizedTitleEncoded := ctx.Params("normalized_title")

	normalizedTitle, err := url.QueryUnescape(normalizedTitleEncoded)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid title encoding"})
	}

	post, err := c.PostService.GetPostByNormalizedTitle(normalizedTitle)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Post not found"})
	}

	return ctx.Status(fiber.StatusOK).JSON(post)
}

// UpdatePost godoc
// @Summary Update a blog post
// @Description Updates a blog post's title or content
// @Tags Posts
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Param request body map[string]string true "Updated Post Data"
// @Success 200 {object} entities.Post
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /admin/posts/{id} [put]
// @security BearerAuth
func (c *PostController) UpdatePost(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var payload struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid payload"})
	}

	post := &entities.Post{
		ID:      id,
		Title:   payload.Title,
		Content: payload.Content,
	}

	if err := c.PostService.UpdatePost(post); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(post)
}

// DeletePost godoc
// @Summary Delete a blog post
// @Description Deletes a blog post by its ID
// @Tags Posts
// @Param id path string true "Post ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /admin/posts/{id} [delete]
// @security BearerAuth
func (c *PostController) DeletePost(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if err := c.PostService.DeletePost(id); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Post deleted successfully"})
}
