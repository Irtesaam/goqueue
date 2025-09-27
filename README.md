<div align="center">
    <h1>GoQueue CLI Application</h1>
</div>

Goqueue is a command-line ToDo app written in Go — because some of us would rather miss deadlines than leave the terminal. It features a modern CLI structure using Cobra and flexible configuration management with Viper.

## Features

### 🎯 Core Functionality
- **Add tasks**: `goq add "Your task description"`
- **List tasks**: `goq list` (or `goq ls`) with table formatting
- **Toggle completion**: `goq toggle 0` to mark tasks as done/undone
- **Edit tasks**: `goq edit 0 "Updated description"`
- **Delete tasks**: `goq delete 0` (or `goq del 0`)
- **Persistent storage**: JSON-based storage with automatic backups

## Installation
### Method 1: Quick Install (Recommended)
The easiest way to install GoQueue:
```bash
curl -fsSL https://raw.githubusercontent.com/Irtesaam/goqueue/master/install.sh | bash
```
This script will:
- Detect your operating system and architecture
- Download the latest pre-built binary
- Install it to your system PATH
- Handle permissions automatically
### Method 2: Build from Source
If you prefer to build from source or the install script doesn't work on your system:
**Prerequisites:** Go 1.19+ and Git
```bash
# Clone and build
git clone https://github.com/Irtesaam/goqueue.git
cd goqueue
make install-user    # Installs to ~/.local/bin (no sudo required)
# OR
make install         # Installs to /usr/local/bin (requires sudo)
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

## Technical Details

### Dependencies
- **Cobra**: Professional CLI framework used by Docker, Kubernetes, etc
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

## Future Enhancements🔥

The current architecture supports easy extension:
- **Priority levels**: Add priority field to todos
- **Categories/Tags**: Organize todos by project or context
- **Due dates**: Add deadline support with notifications
- **Themes**: Customizable colors and formatting
- **Sync**: Cloud synchronization capabilities
- **Import/Export**: Support for other todo formats
- **Search**: Full-text search across todos
- **Statistics**: Completion rates and productivity metrics

## Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature-name`
3. Commit changes: `git commit -am 'Add feature'`
4. Push to branch: `git push origin feature-name`
5. Create Pull Request

## License

[MIT License](LICENSE) - see LICENSE file for details.

<div align="center">
Made with love ❤️ by Irtesaam
</div>
