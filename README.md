# SCP-████: The Self-Modifying Codebase

An interactive CLI game that teaches Git concepts through an SCP Foundation horror/mystery narrative.

## 🎮 About

You are a newly assigned Junior Researcher in the SCP Foundation's Digital Anomalies Division. Your assignment: contain an anomalous self-modifying codebase using proper Git procedures. But be careful - each command you execute teaches the entity more about our reality...

## 🚀 Quick Start

### Installation

```bash
# Clone the repository
git clone https://github.com/byvfx/git-anomaly.git
cd git-anomaly

# Build the game
make build

# Run the game
make run
```

### Alternative: Direct Go Run

```bash
go run main.go
```

## 🎯 How to Play

1. Start the game and type `start` to begin containment protocols
2. Use Git commands to contain the anomaly:
   - `git init` - Initialize containment repository
   - `git add <file>` - Stage files for containment
   - `git commit -m "message"` - Secure files in containment
   - `git status` - Check repository status
   - `git branch` - Manage containment branches
   - `git checkout` - Switch between branches

3. Monitor your stats:
   - **Anomaly Level**: Increases with mistakes (100% = game over)
   - **Researcher Sanity**: Decreases with stress (0% = game over)
   - **Containment Status**: SECURE → BREACH → CRITICAL

4. Take breaks with `tea` command to restore sanity

## 📋 Commands

| Command | Description |
|---------|-------------|
| `help` | Display available commands |
| `start` | Begin the game |
| `status` | Check current game status |
| `tea` | Take a break (restore sanity) |
| `clear` | Clear the screen |
| `quit` | Exit the game |

## 🏗️ Building from Source

### Requirements

- Go 1.21 or higher
- Make (optional, for Makefile commands)

### Build Commands

```bash
# Build for current platform
make build

# Build for all platforms
make build-all

# Run tests
make test

# Format code
make fmt
```

## 📚 Game Progression

### Level 1: Initial Containment
Learn basic Git commands while containing a simple anomalous file.

### Level 2: Staging Area Protocols
Master selective staging to avoid contaminating your repository.

### Level 3: Branch Containment Procedures
Use branches to isolate and contain spreading anomalies.

## 🤝 Contributing

This is a learning project designed to teach Git concepts through interactive storytelling. Contributions are welcome!

## 📜 License

This project is open source and available under the MIT License.

---

**Warning**: This software simulates anomalous digital entities. The SCP Foundation is not responsible for any existential dread, recursive nightmares, or spontaneous Git expertise that may result from playing this game.

*Secure. Contain. Protect.*
