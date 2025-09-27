#!/bin/bash

# GoQueue Installation Script
# This script downloads and installs the latest version of GoQueue

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Configuration
REPO="yourusername/goqueue"
INSTALL_DIR="/usr/local/bin"
USER_INSTALL_DIR="$HOME/.local/bin"
BINARY_NAME="goq"

# Functions
print_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Detect OS and Architecture
detect_platform() {
    local os=$(uname -s | tr '[:upper:]' '[:lower:]')
    local arch=$(uname -m)

    case $os in
        linux)
            OS="linux"
            ;;
        darwin)
            OS="darwin"
            ;;
        *)
            print_error "Unsupported operating system: $os"
            exit 1
            ;;
    esac

    case $arch in
        x86_64|amd64)
            ARCH="amd64"
            ;;
        arm64|aarch64)
            ARCH="arm64"
            ;;
        *)
            print_error "Unsupported architecture: $arch"
            exit 1
            ;;
    esac

    PLATFORM="${OS}-${ARCH}"
    print_info "Detected platform: $PLATFORM"
}

# Check if command exists
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Get latest release version
get_latest_version() {
    if command_exists curl; then
        VERSION=$(curl -s "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
    elif command_exists wget; then
        VERSION=$(wget -qO- "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
    else
        print_error "Neither curl nor wget is available. Please install one of them."
        exit 1
    fi

    if [ -z "$VERSION" ]; then
        print_error "Failed to get latest version"
        exit 1
    fi

    print_info "Latest version: $VERSION"
}

# Download binary
download_binary() {
    local url="https://github.com/$REPO/releases/download/$VERSION/${BINARY_NAME}-${PLATFORM}"
    local tmp_file="/tmp/${BINARY_NAME}"

    print_info "Downloading $BINARY_NAME from $url"

    if command_exists curl; then
        curl -L -o "$tmp_file" "$url"
    elif command_exists wget; then
        wget -O "$tmp_file" "$url"
    fi

    if [ ! -f "$tmp_file" ]; then
        print_error "Failed to download binary"
        exit 1
    fi

    chmod +x "$tmp_file"
    DOWNLOADED_BINARY="$tmp_file"
    print_success "Downloaded successfully"
}

# Install binary
install_binary() {
    # Try system installation first
    if [ -w "$INSTALL_DIR" ] || [ "$EUID" -eq 0 ]; then
        print_info "Installing to $INSTALL_DIR (system-wide)"
        mv "$DOWNLOADED_BINARY" "$INSTALL_DIR/$BINARY_NAME"
        print_success "Installed to $INSTALL_DIR/$BINARY_NAME"
    else
        # Try with sudo
        if command_exists sudo; then
            print_info "Installing to $INSTALL_DIR (requires sudo)"
            sudo mv "$DOWNLOADED_BINARY" "$INSTALL_DIR/$BINARY_NAME"
            print_success "Installed to $INSTALL_DIR/$BINARY_NAME"
        else
            # Fall back to user installation
            print_warning "Cannot install system-wide, installing to user directory"
            mkdir -p "$USER_INSTALL_DIR"
            mv "$DOWNLOADED_BINARY" "$USER_INSTALL_DIR/$BINARY_NAME"
            print_success "Installed to $USER_INSTALL_DIR/$BINARY_NAME"

            # Check if user bin is in PATH
            if [[ ":$PATH:" != *":$USER_INSTALL_DIR:"* ]]; then
                print_warning "⚠️  $USER_INSTALL_DIR is not in your PATH"
                print_info "Add this line to your shell profile (~/.bashrc, ~/.zshrc, etc.):"
                echo "export PATH=\"\$HOME/.local/bin:\$PATH\""
                echo ""
                print_info "Then reload your shell or run: source ~/.bashrc"
            fi
        fi
    fi
}

# Verify installation
verify_installation() {
    if command_exists "$BINARY_NAME"; then
        print_success "✅ GoQueue installed successfully!"
        print_info "Version: $($BINARY_NAME --version 2>/dev/null || echo 'Unknown')"
        echo ""
        print_info "🎉 You can now use GoQueue with the 'goq' command"
        print_info "Try: goq add \"My first todo\""
        print_info "     goq list"
    else
        print_error "Installation verification failed. GoQueue may not be in your PATH."
        print_info "Try running: hash -r"
        print_info "Or start a new terminal session"
    fi
}

# Build from source (fallback)
build_from_source() {
    print_info "Building from source..."

    if ! command_exists git; then
        print_error "Git is not installed. Please install git and try again."
        exit 1
    fi

    if ! command_exists go; then
        print_error "Go is not installed. Please install Go 1.19+ and try again."
        exit 1
    fi

    local tmp_dir="/tmp/goqueue-build"
    rm -rf "$tmp_dir"

    print_info "Cloning repository..."
    git clone "https://github.com/$REPO.git" "$tmp_dir"
    cd "$tmp_dir"

    print_info "Building binary..."
    go build -o "$BINARY_NAME"

    DOWNLOADED_BINARY="$tmp_dir/$BINARY_NAME"
    install_binary

    # Cleanup
    rm -rf "$tmp_dir"
}

# Main installation process
main() {
    echo "🚀 GoQueue Installation Script"
    echo "=============================="
    echo ""

    detect_platform

    # Try to download pre-built binary
    if get_latest_version && download_binary; then
        install_binary
        verify_installation
    else
        print_warning "Failed to download pre-built binary, trying to build from source..."
        build_from_source
        verify_installation
    fi

    echo ""
    print_success "🎉 Installation complete!"
    echo ""
    print_info "📚 Documentation: https://github.com/$REPO"
    print_info "🐛 Issues: https://github.com/$REPO/issues"
    echo ""
    print_info "Get started:"
    echo "  goq add \"Buy groceries\""
    echo "  goq add \"Learn something new\""
    echo "  goq list"
    echo "  goq --help"
}

# Run main function
main "$@"
