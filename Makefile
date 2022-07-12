include .env

swaginit: swag init --parseDependency --parseInternal

start: go run ./main.go