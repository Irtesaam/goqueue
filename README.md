<div align="center">
    <h1>GoQueue CLI Application</h1>
</div>

Goqueue is a command-line To-Do app written in Go — for people who’d rather miss a deadline than leave the terminal. It combines a modern CLI powered by Cobra with flexible configuration via Viper, giving you a fast, no-nonsense workflow right in your terminal.

<img width="1361" height="778" alt="image" src="https://github.com/user-attachments/assets/8180f9fe-54c5-4b3b-bc48-901a1ca43381" />

## Features

### 🎯 Core Functionality
- **Add tasks**: `goq add "Your task description"`
- **List tasks**: `goq list` (or `goq ls`) with table formatting
- **Toggle completion**: `goq toggle 0` to mark tasks as done/undone
- **Edit tasks**: `goq edit 0 "Updated description"`
- **Delete tasks**: `goq delete 0` (or `goq del 0`)
- **Persistent storage**: JSON-based storage with automatic backups

## Installation
### Quick Install (Recommended)
The easiest way to install GoQueue:
```bash
curl -fsSL https://raw.githubusercontent.com/Irtesaam/goqueue/master/install.sh | bash
```
This script will:
- Detect your operating system and architecture
- Download the latest pre-built binary
- Install it to your system PATH
- Handle permissions automatically

### Uninstall GoQueue
To remove GoQueue and all its data:
```bash
sudo rm /usr/local/bin/goq
rm -rf ~/.config/goqueue
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
