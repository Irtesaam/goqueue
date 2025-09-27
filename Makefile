# GoQueue Makefile
# Build and installation automation

.PHONY: build install uninstall clean test help

# Default target
all: build

# Build the binary
build:
	@echo "🔨 Building GoQueue..."
	@echo "📦 Resolving dependencies..."
	go mod tidy
	go build -o goq
	@echo "✅ Build complete! Binary: ./goq"

# Install to system directory (requires sudo)
install: build
	@echo "📦 Installing GoQueue to /usr/local/bin..."
	sudo cp goq /usr/local/bin/goq
	@echo "✅ GoQueue installed successfully!"
	@echo "🎉 You can now use 'goq' from anywhere!"

# Install to user directory (no sudo required)
install-user: build
	@echo "📦 Installing GoQueue to ~/.local/bin..."
	mkdir -p ~/.local/bin
	cp goq ~/.local/bin/goq
	@echo "✅ GoQueue installed to user directory!"
	@echo "ℹ️  Make sure ~/.local/bin is in your PATH"
	@echo "   Add this to your shell profile: export PATH=\"\$$HOME/.local/bin:\$$PATH\""

# Uninstall from system directory
uninstall:
	@echo "🗑️  Uninstalling GoQueue..."
	sudo rm -f /usr/local/bin/goq
	@echo "✅ GoQueue uninstalled successfully!"

# Uninstall from user directory
uninstall-user:
	@echo "🗑️  Uninstalling GoQueue from user directory..."
	rm -f ~/.local/bin/goq
	@echo "✅ GoQueue uninstalled from user directory!"

# Clean build artifacts
clean:
	@echo "🧹 Cleaning build artifacts..."
	rm -f goq
	@echo "✅ Clean complete!"

# Run tests
test:
	@echo "🧪 Running tests..."
	go test ./...
	@echo "✅ Tests complete!"

# Display help
help:
	@echo "GoQueue Build Commands:"
	@echo ""
	@echo "  build        - Build the binary"
	@echo "  install      - Install to /usr/local/bin (requires sudo)"
	@echo "  install-user - Install to ~/.local/bin (no sudo required)"
	@echo "  uninstall    - Remove from /usr/local/bin"
	@echo "  uninstall-user - Remove from ~/.local/bin"
	@echo "  clean        - Remove build artifacts"
	@echo "  test         - Run tests"
	@echo "  help         - Show this help message"
	@echo ""
	@echo "Examples:"
	@echo "  make build"
	@echo "  make install-user"
	@echo "  make uninstall"
