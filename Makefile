.PHONY: all clean build run tidy

APP_NAME=watcher

all: run

clean:
	@echo "Cleaning up..."
	@if [ -d "build" ]; then rm -rf build; fi
	@if [ -d "dist" ]; then rm -rf dist; fi

build: clean
	@echo "Building..."
	@mkdir -p build

	@echo "Compiling..."
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/$(APP_NAME) cmd/watcher/main.go

tidy:
	@echo "Tidying up..."
	@go mod tidy

run:tidy build 
	@echo "Running..." 
	@cd build;./$(APP_NAME) connect

upx:
	@echo "Compressing..."
	@upx -9 build/$(APP_NAME)

release: build upx
	@echo "Releasing..."
	@mkdir -p dist
	@cp build/$(APP_NAME) dist
	@cd dist;tar -czvf $(APP_NAME).tar.gz $(APP_NAME)
	@rm -rf dist/$(APP_NAME)
	@echo "Release done! dist/$(APP_NAME).tar.gz"
	@mv dist/$(APP_NAME).tar.gz /mnt/d/