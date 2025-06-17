run: build
	@./bin/app

templ:
	@templ generate --watch --proxy=http://localhost:3000 --open-browser=false

tailwind:
	npx tailwindcss -i views/css/app.css -o public/styles.css --watch

install:
	@echo "Installing dependencies..."
	@go install github.com/a-h/templ/cmd/templ@latest
	@curl -sLo tailwindcss https://github.com/tailwindlabs/tailwindcss/releases/download/v3.3.5/tailwindcss-linux-x64
	@chmod +x tailwindcss
	@mv tailwindcss /usr/local/bin/
	@go get ./...
	@go mod vendor
	@go mod tidy
	@go mod download
	@tailwindcss -i views/css/app.css -o public/styles.css --watch
	@go build -o bin/app main.go


build: 
	@tailwindcss -i views/css/app.css -o public/styles.css --watch
	@go build -o bin/app main.go

css:
	npx tailwindcss -i views/css/app.css -o public/styles.css --watch  
