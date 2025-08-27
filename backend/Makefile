.PHONY: help build run clean test dev

# デフォルトターゲット
help:
	@echo "Available commands:"
	@echo "  build     - Build the Go application"
	@echo "  run       - Run the Go application"
	@echo "  dev       - Start development mode"
	@echo "  test      - Run tests"
	@echo "  clean     - Clean build artifacts"

# Go アプリケーションをビルド
build:
	go build -o bin/go-sse .

# Go アプリケーションを実行
run: build
	./bin/go-sse

# 開発モードでサーバーを起動
dev:
	go run .

# テスト実行
test:
	go test ./...

# クリーンアップ
clean:
	rm -rf bin/

# 依存関係のインストール
install:
	go mod tidy