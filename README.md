# Professional GoQueue CLI Application

A beautiful and professional command-line todo application built with Go, featuring a modern CLI structure using Cobra and flexible configuration management with Viper.

## Features

### ✨ Professional CLI Structure with Cobra
- **Subcommands**: `goq add`, `goq list`, `goq edit`, etc. (instead of flags like `--add`)
- **Built-in help**: Comprehensive help system with `--help` for each command
- **Command aliases**: Use `goq ls` instead of `goq list`, `goq del` instead of `goq delete`
- **Input validation**: Proper argument validation and error messages
- **Modern UX**: Following industry standards used by tools like Docker, Kubernetes, and Git

### ⚙️ Robust Configuration with Viper
- **Configuration file**: `~/.config/goqueue/config.yaml`
- **Environment variables**: Override settings with `GOQUEUE_FILE=/path/to/file`
- **Command-line flags**: Use `--file` to specify custom file location
- **Automatic directory creation**: Config directory created automatically
- **Priority order**: CLI flags > Environment variables > Config file > Defaults

### 🎯 Core Todo Functionality
- **Add tasks**: `goq add "Your task description"`
- **List tasks**: `goq list` (or `goq ls`) with beautiful table formatting
- **Toggle completion**: `goq toggle 0` to mark tasks as done/undone
- **Edit tasks**: `goq edit 0 "Updated description"`
- **Delete tasks**: `goq delete 0` (or `goq del 0`)
- **Persistent storage**: JSON-based storage with automatic backups

## Installation

### Quick Install (Recommended)

#### Option 1: Download and Install Script
```bash
# Download and run the install script
curl -fsSL https://raw.githubusercontent.com/Irtesaam/goqueue/main/install.sh | bash
```

#### Option 2: Manual Installation
```bash
# Clone the repository
git clone https://github.com/Irtesaam/goqueue.git
cd goqueue

# Build and install
make install
```

#### Option 3: Download Pre-built Binary
```bash
# Download the latest release for your platform
# Linux (x86_64)
curl -L -o goq https://github.com/Irtesaam/goqueue/releases/latest/download/goq-linux-amd64
chmod +x goq
sudo mv goq /usr/local/bin/

# macOS (x86_64)
curl -L -o goq https://github.com/Irtesaam/goqueue/releases/latest/download/goq-darwin-amd64
chmod +x goq
sudo mv goq /usr/local/bin/

# macOS (ARM64 - Apple Silicon)
curl -L -o goq https://github.com/Irtesaam/goqueue/releases/latest/download/goq-darwin-arm64
chmod +x goq
sudo mv goq /usr/local/bin/
```

### Build from Source

#### Prerequisites
- Go 1.19 or later installed on your system
- Git for cloning the repository

#### Steps
```bash
# Clone the repository
git clone https://github.com/Irtesaam/goqueue.git
cd goqueue

# Resolve dependencies and build
go mod tidy
go build -o goq

# Install globally (optional)
sudo mv goq /usr/local/bin/

# Or install to user directory (no sudo required)
mkdir -p ~/.local/bin
mv goq ~/.local/bin/
# Make sure ~/.local/bin is in your PATH
echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.bashrc  # or ~/.zshrc
source ~/.bashrc  # or source ~/.zshrc
```

### Verify Installation
```bash
# Check if goq is installed correctly
goq --version
goq --help

# Add your first todo
goq add "Welcome to GoQueue! 🎉"
goq list
```

## Usage

### Basic Commands

```bash
# Add a new todo
goq add "Buy groceries"
goq add "Complete the project report"

# List all todos
goq list           # or: goq ls, goq show

# Toggle completion status
goq toggle 0       # Toggle first task

# Edit a task
goq edit 1 "Updated task description"

# Delete a task
goq delete 2       # or: goq del 2, goq rm 2

# Get help
goq --help
goq add --help     # Help for specific commands
```

### Configuration Examples

#### 1. Using configuration file
Create `~/.config/goqueue/config.yaml`:
```yaml
file: "/path/to/your/custom/todos.json"
```

#### 2. Using environment variables
```bash
export GOQUEUE_FILE="/tmp/work-todos.json"
goq add "Work-related task"
```

#### 3. Using command-line flags
```bash
goq --file "/tmp/personal-todos.json" add "Personal task"
goq --file "/tmp/personal-todos.json" list
```

#### 4. Multiple todo lists
```bash
# Work todos
alias work-goq='goq --file "$HOME/.config/goqueue/work.json"'
work-goq add "Review code"

# Personal todos
alias personal-goq='goq --file "$HOME/.config/goqueue/personal.json"'
personal-goq add "Buy milk"
```

## Architecture

### Project Structure
```
.
├── cmd/                      # Cobra CLI commands
│   ├── add.go              # Add todo command
│   ├── delete.go           # Delete todo command
│   ├── edit.go             # Edit todo command
│   ├── list.go             # List todos command
│   ├── root.go             # Root command and Viper config
│   ├── toggle.go           # Toggle completion command
│   └── version.go          # Version command
├── .github/
│   └── workflows/
│       └── release.yml     # GitHub Actions for automated releases
├── internal/               # Internal packages (not importable by external projects)
│   ├── storage/            # Storage abstraction layer
│   │   └── storage.go      # Generic JSON storage implementation
│   └── todo/               # Todo business logic
│       └── todo.go         # Todo types and methods
├── .gitignore              # Git ignore rules
├── config.example.yaml     # Example configuration file
├── go.mod                  # Go module definition
├── go.sum                  # Go module checksums
├── install.sh              # Installation script for users
├── LICENSE                 # MIT license
├── main.go                 # Application entry point
├── Makefile                # Build and installation automation
└── README.md               # Project documentation
```

### Key Design Patterns
- **Command Pattern**: Each CLI command is a separate, focused module
- **Dependency Injection**: Storage and configuration injected into commands
- **Generic Storage**: Type-safe JSON storage that can be reused
- **Configuration Hierarchy**: Multiple configuration sources with clear precedence
- **Clean Architecture**: Business logic separated from CLI concerns

## Migration from Basic CLI

The application has been upgraded from a basic flag-based CLI to a professional subcommand structure:

### Before (Old)
```bash
./todo -add "Task"
./todo -list
./todo -toggle 0
./todo -del 1
```

### After (New)
```bash
goq add "Task"
goq list
goq toggle 0
goq delete 1
```

## Technical Details

### Dependencies
- **Cobra**: Professional CLI framework used by Docker, Kubernetes, Hugo
- **Viper**: Configuration management supporting files, env vars, and flags
- **Aquasecurity Table**: Beautiful table formatting for list output

### Configuration Precedence
1. Command-line flags (`--file`)
2. Environment variables (`GOQUEUE_FILE`)
3. Configuration file (`~/.config/goqueue/config.yaml`)
4. Default values (`~/.config/goqueue/todos.json`)

### Storage Format
Tasks are stored in JSON format:
```json
[
    {
        "Title": "Learn Cobra CLI framework",
        "Completed": true,
        "CreatedAt": "2025-09-27T09:24:12.123456789+05:30",
        "CompletedAt": "2025-09-27T09:24:31.987654321+05:30"
    }
]
```

## Future Enhancements

The current architecture supports easy extension:
- **Priority levels**: Add priority field to todos
- **Categories/Tags**: Organize todos by project or context
- **Due dates**: Add deadline support with notifications
- **Themes**: Customizable colors and formatting
- **Sync**: Cloud synchronization capabilities
- **Import/Export**: Support for other todo formats
- **Search**: Full-text search across todos
- **Statistics**: Completion rates and productivity metrics

## For Developers

### Setting up the Project on GitHub

1. **Create a new repository on GitHub** named `goqueue`

2. **Initialize and push your code:**
```bash
# Initialize git (if not already done)
git init

# Add all files
git add .

# Commit
git commit -m "Initial commit: Professional GoQueue CLI with Cobra and Viper"

# Add remote (replace 'Irtesaam' with your GitHub username)
git remote add origin https://github.com/Irtesaam/goqueue.git

# Push to GitHub
git push -u origin main
```

3. **Enable GitHub Actions** (automatically enabled when you push the workflow file)

4. **Create your first release:**
```bash
# Tag a version
git tag v1.0.0
git push origin v1.0.0
```

This will automatically trigger the GitHub Actions workflow to build binaries for all platforms and create a release.

### Local Development

```bash
# Clone your repository
git clone https://github.com/Irtesaam/goqueue.git
cd goqueue

# Build and test locally
make build
./goq --help

# Install locally for testing
make install-user

# Run from anywhere
goq add "Test from anywhere!"
goq list
```

### Release Process

1. **Update version** and commit changes
2. **Create and push a git tag:**
```bash
git tag v1.0.1
git push origin v1.0.1
```
3. **GitHub Actions will automatically:**
   - Build binaries for Linux, macOS, and Windows
   - Create a GitHub release
   - Upload all binaries as release assets
   - Generate release notes

### Testing the Installation Process

After pushing to GitHub, test the installation process:

```bash
# Test the install script
curl -fsSL https://raw.githubusercontent.com/Irtesaam/goqueue/main/install.sh | bash

# Test manual installation
curl -L -o goq https://github.com/Irtesaam/goqueue/releases/latest/download/goq-linux-amd64
chmod +x goq
./goq --version
```

## Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature-name`
3. Commit changes: `git commit -am 'Add feature'`
4. Push to branch: `git push origin feature-name`
5. Create Pull Request

### Development Guidelines

- Follow Go best practices and idioms
- Add tests for new functionality
- Update documentation for new features
- Use conventional commit messages
- Ensure cross-platform compatibility

## License

MIT License - see LICENSE file for details.
