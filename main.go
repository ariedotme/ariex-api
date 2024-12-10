package main

import (
	"log"
	"os"

	"github.com/ariedotme/ariex-backend/application/services"
	"github.com/ariedotme/ariex-backend/application/usecases"
	_ "github.com/ariedotme/ariex-backend/docs"
	"github.com/ariedotme/ariex-backend/infrastructure/database"
	"github.com/ariedotme/ariex-backend/infrastructure/http"
	"github.com/ariedotme/ariex-backend/infrastructure/http/controllers"
	"github.com/ariedotme/ariex-backend/shared/env"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title Ariex API
// @version 1.0
// @description API for managing blog posts, users, and contact messages
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @contact.name Aurora Surreaux
// @contact.url https://github.com/ariedotme/
// @contact.email contatoaurorash@gmail.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:3000
func main() {
	env.LoadEnv()

	database.InitSupabase()

	userRepository := &database.UserRepositoryImpl{}
	postRepository := &database.PostRepositoryImpl{}

	postService := &services.PostService{PostRepository: postRepository}
	userService := &services.UserService{UserRepository: userRepository}
	emailService := &services.EmailService{
		SMTPHost: os.Getenv("SMTP_HOST"),
		SMTPPort: os.Getenv("SMTP_PORT"),
		Username: os.Getenv("SMTP_USERNAME"),
		Password: os.Getenv("SMPT_PASSWORD"),
	}

	loginUserUseCase := &usecases.LoginUserUseCase{UserService: userService}
	createPostUseCase := &usecases.CreatePostUseCase{PostService: postService}
	sendContactUseCase := &usecases.SendContactUseCase{EmailService: emailService}

	userController := &controllers.UserController{LoginUserUseCase: loginUserUseCase, UserService: userService}
	postController := &controllers.PostController{CreatePostUseCase: createPostUseCase, PostService: postService}
	contactController := &controllers.ContactController{SendContactUseCase: sendContactUseCase}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("FRONTEND_URL"),
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Authorization",
	}))

	http.SetupRoutes(app, userController, postController, contactController)

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	log.Fatal(app.Listen(":4000"))
}
