package http

import (
	"github.com/ariedotme/ariex-backend/infrastructure/http/controllers"
	"github.com/ariedotme/ariex-backend/infrastructure/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, userController *controllers.UserController, postController *controllers.PostController, contactController *controllers.ContactController) {
	app.Use(middlewares.LoggingMiddleware)

	app.Post("/register", middlewares.BasicAuthMiddleware, userController.RegisterUser)

	auth := app.Group("/auth")
	auth.Post("/login", userController.Login)

	// public routes
	app.Get("/posts", postController.GetPosts)
	app.Get("/posts/:id", postController.GetPostByID)
	app.Get("/posts/title/:normalized_title", postController.GetPostByNormalizedTitle)
	app.Post("/contact", contactController.SendContact)

	admin := app.Group("/admin", middlewares.AuthMiddleware)
	admin.Post("/posts", postController.CreatePost)
	admin.Put("/posts/:id", postController.UpdatePost)
	admin.Delete("/posts/:id", postController.DeletePost)
}
