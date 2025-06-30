# SCP-████: The Self-Modifying Codebase

An interactive CLI game that teaches Git concepts through an SCP Foundation horror/mystery narrative.

## 🎮 About

You are a newly assigned Junior Researcher in the SCP Foundation's Digital Anomalies Division. Your assignment: contain an anomalous self-modifying codebase using proper Git procedures. But be careful - each command you execute teaches the entity more about our reality...

> ⚠️ **Early Development Notice**: This game is currently in early development. Features, gameplay mechanics, and story elements are subject to change. Some planned features may not yet be implemented.

## 🚀 Quick Start

### Installation

```bash
# Clone the repository
git clone https://github.com/byvfx/git-anomaly.git
cd git-anomaly

# Build the game
make build

# Run the game (modern TUI interface)
make run

# Or run with classic CLI interface
./bin/scp-git --classic
```

### Alternative: Direct Go Run

```bash
# Modern TUI interface (default)
go run main.go

# Classic CLI interface
go run main.go --classic
```

## 🎯 How to Play

> **📋 Note**: The modern TUI interface is temporarily on hold while we focus on perfecting the core CLI experience. The classic CLI now provides the complete game with all enhanced features including tab completion and command history.

### Enhanced CLI Interface (Current)
Experience the complete game through our improved command-line interface featuring:
- 🚀 **Tab Completion**: Smart completion for all git commands, filenames, and branch names
- 📜 **Command History**: Navigate previous commands with up/down arrows
- 🎓 **Official Git Tutorial**: 4 progressive levels following proven Git learning structure
- 🎨 **SCP Theming**: Full Foundation aesthetic with color-coded output
- 🔧 **Advanced Git Commands**: Complete git workflow from config to merge

### Gameplay
1. Type `start` to begin containment protocols
2. Follow the progressive 4-level Git tutorial:

#### Level 1: Initial Containment Setup
   - `git config user.name "Your Name"` - Configure researcher identity
   - `git config user.email "email@site.com"` - Set contact protocol
   - `git init` - Initialize containment repository
   - `git add .` - Stage all discovered files
   - `git commit -m "Initial containment"` - Secure files

#### Level 2: Monitoring Changes
   - `git status` - Check for autonomous modifications
   - `git diff` - Analyze specific changes made by the entity
   - `git commit -a -m "Document changes"` - Commit all tracked modifications

#### Level 3: Historical Analysis
   - `git log` - Review complete containment timeline
   - `git log -p` - Examine detailed change history
   - `git show <commit>` - Investigate specific incidents

#### Level 4: Parallel Containment Strategies
   - `git switch -c strategy-a` - Create experimental branch
   - `git merge strategy-a` - Integrate successful approaches

3. **Enhanced Features**:
   - Use **Tab** for command completion
   - Use **Up/Down arrows** for command history
   - Type `brief` to review current level objectives
   - Type `help` for complete command reference

4. Monitor your stats:
   - **Anomaly Level**: Increases with mistakes (100% = game over)
   - **Researcher Sanity**: Decreases with stress (0% = game over)  
   - **Containment Status**: SECURE → BREACH → CRITICAL

## 📋 Commands

### System Commands
| Command | Description |
|---------|-------------|
| `help` | Display available commands |
| `start` | Begin the game |
| `status` | Check current game status |
| `brief` / `briefing` | Re-display current level objectives |
| `breathe` | Take a moment to recover sanity (+5%) |
| `clear` | Clear the screen |
| `quit` | Exit the game |

### Git Commands  
| Command | Description |
|---------|-------------|
| `git config user.name "Name"` | Configure researcher identity |
| `git config user.email "email"` | Set Foundation contact |
| `git init` | Initialize containment repository |
| `git add <file>` | Stage files for containment |
| `git add .` | Stage all files |
| `git commit -m "msg"` | Secure files in containment |
| `git commit -a -m "msg"` | Commit all tracked changes |
| `git status` | View repository status |
| `git diff` | Show file modifications |
| `git log` | View containment history |
| `git log -p` | View history with changes |
| `git show [commit]` | Examine specific commit |
| `git branch [name]` | Create or list branches |
| `git switch <branch>` | Switch to existing branch |
| `git switch -c <branch>` | Create and switch to new branch |
| `git checkout <branch>` | Switch branches (classic) |
| `git merge <branch>` | Merge containment strategies |

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

Following the official Git tutorial structure for optimal learning:

### Level 1: Initial Containment Setup
Configure your researcher identity and establish version control for all discovered files. Learn the fundamentals of Git initialization and committing.

### Level 2: Monitoring Changes  
Track autonomous entity modifications using status and diff commands. Learn to document and commit changes as they occur.

### Level 3: Historical Analysis
Investigate the entity's behavior patterns through commit history. Master log analysis and forensic examination techniques.

### Level 4: Parallel Containment Strategies
Implement multiple containment approaches using branches, then merge successful strategies. Advanced workflow management.

## 🤝 Contributing

This is a learning project designed to teach Git concepts through interactive storytelling. Contributions are welcome!

## 📜 License

This project is open source and available under the MIT License.

---

**Warning**: This software simulates anomalous digital entities. The SCP Foundation is not responsible for any existential dread, recursive nightmares, or spontaneous Git expertise that may result from playing this game.

*Secure. Contain. Protect.*
