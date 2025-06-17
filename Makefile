run: build
	@./bin/app

templ:
	@templ generate --watch --proxy=http://localhost:3000 --open-browser=false

tailwind:
	npx tailwindcss -i views/css/app.css -o public/styles.css --watch

install:
	@echo "Installing dependencies..."
	@go install github.com/a-h/templ/cmd/templ@latest
	@go install github.com/tailwindlabs/tailwindcss@latest
	@go get ./...
	@go mod vendor
	@go mod tidy
	@go mod download

build: 
	@tailwindcss -i views/css/app.css -o public/styles.css --watch
	@go build -o bin/app main.go

css:
	npx tailwindcss -i views/css/app.css -o public/styles.css --watch  
