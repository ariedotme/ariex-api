# Ariex API

### I won't develop it any further, I migrated to a Elixir backend.


This is the API for my personal site, developed in Golang with fiber and following a DDD Lite architectural approach.

Use `swag init --output ./docs` before launching the application (`go run main.go`) to generate a Swagger configuration for the API.

Currently the API handles:

- JWT User Auth (for admin access)
- Blog posts
- Contact through SMTP

On the future the front-end will also be avaiable.
