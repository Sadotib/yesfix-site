run: build
	@./bin/app

templ:
	@templ generate --watch --proxy=http://localhost:4000 --open-browser=false

tailwind:
	npx tailwindcss -i views/css/app.css -o public/styles.css --watch

install:
	@echo "Installing dependencies..."
	@go install github.com/a-h/templ/cmd/templ@latest
	@curl -sLo tailwindcss https://github.com/tailwindlabs/tailwindcss/releases/download/v3.3.5/tailwindcss-linux-x64
	@chmod +x tailwindcss
	@mkdir -p bin
	@mv tailwindcss bin/
	@echo "Running Tailwind..."
	@./bin/tailwindcss -i views/css/app.css -o public/styles.css --minify || { echo "Tailwind build failed"; exit 1; }
	@go get ./...
	@go mod vendor
	@go mod tidy
	@go mod download
	@go build -o bin/app main.go



build: 
	@npx tailwindcss -i views/css/app.css -o public/styles.css
	@go build -o bin/app main.go

css:
	npx tailwindcss -i views/css/app.css -o public/styles.css --watch

migrate:
	@go install github.com/pressly/goose/v3/cmd/goose@latest
	@goose postgres "$$DATABASE_URL" up
	@go run github.com/pressly/goose/v3/cmd/goose@latest -dir dbConfig/migrationsTest postgres "$$DATABASE_URL" up
	
