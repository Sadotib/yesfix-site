run: build
	@./bin/app

templ:
	@templ generate --watch --proxy=http://localhost:3000 --open-browser=false

tailwind:
	npx tailwindcss -i views/css/app.css -o public/styles.css --watch

install:
	@go install github.com/templ/templ@latest
	@go get ./...
	@go mod vendor
	@go mod tidy
	@go mod download
	@npm install -D tailwindcss

build: 
	@go build -o bin/app .
	@npx tailwindcss -i views/css/app.css -o public/styles.css --watch
	@go build -o bin/app main.go

css:
	npx tailwindcss -i views/css/app.css -o public/styles.css --watch  
