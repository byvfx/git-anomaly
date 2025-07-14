# SCP-████: The Self-Modifying Codebase
## Interactive Git Learning Game Project Plan

### Project Overview
A CLI-based interactive game that teaches Git concepts through an SCP Foundation horror/mystery theme. Players take on the role of a Foundation researcher tasked with containing an anomalous self-modifying codebase using proper Git procedures.

### Technical Stack
- **Language**: Go
- **Primary Dependencies**: 
  - `cobra` (CLI framework)
  - `readline` (classic CLI input handling)
  - `color` (terminal colors)
- **Target**: Cross-platform CLI application

Educational Framework (Based on Official Git Tutorial)
Our SCP containment training follows the proven structure from the official Git documentation:
Level 1: Initial Containment Setup
Objective: Establish basic containment protocols

Git Tutorial Section: "Importing a new project"
Commands: git config, git init, git add ., git commit
SCP Context: Setting up initial containment for SCP-████

Level 2: Monitoring Changes
Objective: Track anomalous modifications to the codebase

Git Tutorial Section: "Making changes"
Commands: git add, git status, git diff, git commit -a
SCP Context: The entity is modifying files - document all changes

Level 3: Historical Analysis
Objective: Investigate the entity's past behavior

Git Tutorial Section: "Viewing project history"
Commands: git log, git log -p, git log --stat, git show
SCP Context: Analyzing timeline of anomalous activity

Level 4: Parallel Containment Strategies
Objective: Test multiple containment approaches simultaneously

Git Tutorial Section: "Managing branches"
Commands: git branch, git switch, git merge, git branch -d
SCP Context: Running parallel experiments to contain the entity

Level 5: Multi-Site Collaboration
Objective: Coordinate with other Foundation sites

Git Tutorial Section: "Using Git for collaboration"
Commands: git clone, git pull, git remote, git fetch
SCP Context: Sharing containment data between Foundation facilities

Level 6: Advanced Investigation Techniques
Objective: Deep forensic analysis of the anomaly

Git Tutorial Section: "Exploring history"
Commands: git grep, git tag, git diff, git reset, commit references (HEAD^, etc.)
SCP Context: Advanced analysis to understand the entity's true nature
---

## Release Management Strategy

### Release Notes Template (`release.md`)

The `release.md` file serves as the current release staging area and follows this structure:

```markdown
# Release Notes

All release notes have been organized by version for easier navigation.

## Current Release

- [v0.1.0 - Foundation Establishment](docs/releases/v0.1.0.md) - Initial SCP containment protocols with basic Git simulation

## Previous Releases

- [v0.0.1 - Initial Containment](docs/releases/v0.0.1.md) - First containment attempt of SCP-████

## Release Highlights

### Latest Features (v0.3.0)
- **Classic CLI Interface**: Streamlined command-line interface with readline support
- **Interactive Git Simulation**: Core git commands (init, add, commit, status, branch)
- **Foundation Documentation**: SCP-style formatting and error messages
- **Anomaly Detection**: Simplified anomaly level tracking and containment status
- **Tab Completion**: Smart autocomplete for git commands and file names

### Core Features
- Foundation-standard CLI interface with proper SCP documentation
- Interactive Git command learning with real-time feedback
- Anomaly level tracking and containment status monitoring
- ASCII art integration with SCP visual theming
- High performance with concurrent anomaly processing

## Installation

**🔒 Foundation Authorized Installation Methods**

**Quick Install (Recommended)**

**Windows (PowerShell):**
```powershell
iwr -useb https://raw.githubusercontent.com/byvfx/scp-git-game/main/scripts/install-windows.ps1 | iex
```

**Linux/macOS:**
```bash
curl -sSL https://raw.githubusercontent.com/byvfx/scp-git-game/main/scripts/install-unix.sh | bash
```

**Manual Installation**
Download from GitHub Releases and extract to your PATH.

**Build from Source**
```bash
git clone https://github.com/byvfx/scp-git-game.git
cd scp-git-game
make build
```

**Detailed installation guide**: [docs/guides/INSTALLATION.md](docs/guides/INSTALLATION.md)
```

### GitHub Actions Integration

The workflow will:
1. Read `release.md` for release notes
2. Build cross-platform binaries
3. Create GitHub release with notes from `release.md`
4. Archive the release notes to `docs/releases/vX.X.X.md`
5. Reset `release.md` template for next cycle

### Versioning Strategy
- **v0.x.x**: Development/Alpha releases (Phase 1-2)
- **v1.0.0**: First stable release (Phase 3 complete)
- **v1.x.x**: Feature releases and improvements
- **vX.Y.Z**: Semantic versioning (Major.Minor.Patch)

---

## Development Phases

### Phase 1: Foundation Setup
**Goal**: Basic CLI structure and core game loop

#### 1.1 Project Structure
```
scp-git-game/
├── cmd/
│   └── root.go           # Main CLI command
├── pkg/
│   ├── game/
│   │   ├── state.go      # Game state management
│   │   ├── levels.go     # Level definitions
│   │   └── commands.go   # Git command simulation
│   ├── scp/
│   │   ├── document.go   # SCP documentation formatter
│   │   └── anomaly.go    # Anomaly effects
│   ├── ui/
│   │   ├── terminal.go   # Terminal output formatting
│   │   └── prompts.go    # User input handling
│   └── config/
│       └── settings.go   # Configuration management
├── docs/
│   ├── guides/
│   │   ├── INSTALLATION.md
│   │   ├── QUICK_START.md
│   │   ├── README.md
│   │   └── SECURITY.md
│   └── releases/
│       ├── v0.1.0.md
│       ├── v0.2.0.md
│       └── ...
├── assets/
│   ├── ascii/            # ASCII art files
│   └── text/             # SCP documents, flavor text
├── scripts/
│   ├── install-windows.ps1  # Windows installation script
│   ├── install-unix.sh      # Linux/macOS installation script
│   └── build.sh             # Build script
├── .github/
│   └── workflows/
│       ├── release.yml      # GitHub Actions for releases
│       └── ci.yml           # Continuous integration
├── Makefile                 # Build and development commands
├── go.mod
├── go.sum
├── README.md
├── release.md               # Current release notes for automation
├── ROADMAP.md
├── CONTRIBUTING.md
└── claude.md                # This file
```

#### 1.2 Core Components
- [ ] Basic CLI with `cobra`
- [ ] Game state struct (repository, staging, branches, etc.)
- [ ] Simple command parser for git commands
- [ ] Basic level progression system
- [ ] SCP-style text formatter

#### 1.3 Testing Standards
- Unit tests for all game logic functions
- Integration tests for command processing
- Mock user input for automated testing
- Test coverage minimum: 80%

#### 1.4 Acceptance Criteria
- [ ] Application starts with SCP-style intro
- [ ] Can execute basic `git init`, `git add`, `git commit`
- [ ] Game state updates correctly
- [ ] Level progression works
- [ ] Clean error handling and user feedback

---

### Phase 2: Enhanced Simulation & Theming
**Goal**: Realistic Git simulation with full SCP theming

#### 2.1 Advanced Git Simulation
- [ ] Complete Git command set (branch, merge, checkout, etc.)
- [ ] Proper staging area simulation
- [ ] Branch management and visualization
- [ ] Commit history with SCP-style messages
- [ ] Merge conflict scenarios

#### 2.2 SCP Theming Implementation
- [ ] SCP document generator (incident reports, containment procedures)
- [ ] Anomaly level system (increases with mistakes)
- [ ] "Researcher notes" as help system
- [ ] Containment breach alerts for errors
- [ ] ASCII art integration

#### 2.3 Containment Protocol Selection
- [ ] **Simulation Protocol** - Safe contained environment (Phase 1 implementation)
- [ ] **Live Protocol** - Real Git integration for advanced researchers (Phase 2 addition)
- [ ] Protocol selection menu with SCP-style warnings and clearance levels
- [ ] Seamless switching between protocols mid-game
- [ ] Progress tracking across both protocols

#### 2.4 Testing Standards
- Integration tests for all Git commands in both protocols
- Anomaly system testing
- Educational effectiveness testing (do users actually learn Git?)
- User experience testing with sample scenarios
- Performance testing for terminal animations
- Real Git integration testing for Live Protocol

#### 2.5 Acceptance Criteria
- [ ] All major Git commands work correctly in both protocols
- [ ] SCP theming is consistent throughout
- [ ] Containment protocol selection enhances immersion
- [ ] Educational progression follows official Git tutorial structure
- [ ] Smooth user experience with proper feedback

---

### Phase 3: Advanced Features & Polish
**Goal**: Full gameplay experience with advanced features

#### 3.1 Advanced Gameplay
- [ ] Multiple SCP scenarios/levels
- [ ] Trapped researcher rescue mechanics
- [ ] Reality distortion effects in terminal
- [ ] Achievement system (Foundation commendations)
- [ ] Save/load game progress

#### 3.2 Terminal Effects
- [ ] Animated ASCII art
- [ ] Glitch effects when anomaly level is high
- [ ] Colored output with SCP color scheme
- [ ] Progress bars styled as containment meters
- [ ] Dynamic terminal resizing support

#### 3.3 Content Creation
- [ ] Complete SCP documentation for the main anomaly
- [ ] Multiple mini-SCPs for different levels
- [ ] Comprehensive help system disguised as Foundation training materials
- [ ] Easter eggs and hidden content

#### 3.4 Testing Standards
- End-to-end gameplay testing
- Cross-platform compatibility testing
- Accessibility testing (color blind support, etc.)
- Performance testing under various terminal conditions
- User acceptance testing with Git beginners

#### 3.5 Acceptance Criteria
- [ ] Complete game experience from start to finish
- [ ] All visual effects work smoothly
- [ ] Game is learnable by Git beginners
- [ ] No crashes or data loss
- [ ] Professional polish level

---

### Phase 4: Distribution & Documentation
**Goal**: Production-ready release with automated release management

#### 4.1 Release Management
- [ ] `release.md` template for GitHub Actions integration
- [ ] Automated release notes generation from `release.md`
- [ ] Semantic versioning strategy
- [ ] GitHub Actions workflow for automated releases
- [ ] Cross-platform binaries (Windows, macOS, Linux)
- [ ] Installation scripts for different platforms

#### 4.2 Documentation Structure
- [ ] Comprehensive guides in `docs/guides/`:
  - [ ] `INSTALLATION.md` - Platform-specific installation instructions
  - [ ] `QUICK_START.md` - Getting started guide
  - [ ] `README.md` - Main documentation
  - [ ] `SECURITY.md` - Security considerations (fitting for SCP theme!)
- [ ] Release documentation in `docs/releases/`:
  - [ ] Individual release notes for each version
  - [ ] Release highlights and feature summaries
  - [ ] Breaking changes documentation
- [ ] Developer documentation
- [ ] Contribution guidelines
- [ ] Educational Git reference guide integrated with SCP lore

#### 4.3 GitHub Actions Integration
```yaml
# .github/workflows/release.yml
# Automated release workflow that:
# - Uses release.md as release notes
# - Builds cross-platform binaries
# - Creates GitHub releases
# - Updates documentation
---


## Testing Strategy

### Unit Testing
```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Run with race detection
go test -race ./...
```

### Integration Testing
- Test complete user journeys through levels
- Mock terminal interactions for automated testing
- Verify game state persistence

### Manual Testing Checklist
- [ ] All Git commands produce expected results
- [ ] Error messages are helpful and in-character
- [ ] Level progression is logical and educational
- [ ] Terminal effects don't break on different systems
- [ ] Game is actually fun and engaging

### Performance Testing
- Memory usage during long gameplay sessions
- Terminal rendering performance
- Startup time optimization

---

## Success Metrics

### Educational Effectiveness
- Players can successfully use basic Git commands after playing
- Progression through levels correlates with Git knowledge
- Error recovery teaches proper Git practices

### Engagement Metrics
- Average play session length
- Level completion rates
- User feedback on SCP theming effectiveness

### Technical Quality
- Zero crashes during normal gameplay
- Sub-second response time for all commands
- Clean installation on all target platforms

---

## Development Workflow

### Development Environment Setup
```bash
# Initialize project
go mod init scp-git-game

# Install dependencies
go get github.com/spf13/cobra@latest
go get github.com/charmbracelet/bubbletea@latest
go get github.com/charmbracelet/lipgloss@latest
go get github.com/fatih/color@latest

# Development tools
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Build the project
make build

# Or manually:
go build -o bin/scp-git ./cmd
```

### Code Quality Standards
- Follow Go best practices and idioms
- Use `golangci-lint` for static analysis
- Maintain consistent code formatting with `gofmt`
- Write meaningful commit messages (ironically important for a Git game!)

### Git Workflow
- Feature branches for each phase
- Pull requests with thorough testing
- Semantic versioning for releases
- Detailed commit messages explaining changes

---

## Notes for Claude Code

### When Testing Components:
1. Start with basic functionality before adding theming
2. Test each Git command individually before combining
3. Verify error handling works correctly
4. Check terminal output formatting on different screen sizes

### Debugging Tips:
- Use `fmt.Printf` for debugging game state
- Test with various terminal types (bash, zsh, PowerShell)
- Verify ASCII art displays correctly
- Check color output on different terminals

### File Organization:
- Keep SCP documents in separate text files for easy editing
- Use embedded files for ASCII art
- Separate game logic from presentation logic
- Make components testable in isolation

---

## MVP Definition

### Minimum Viable Product for Phase 1:

- Working CLI with Simulation Protocol that teaches the first 3 levels of Git commands
- Commands: git config, git init, git add, git commit, git status, git diff, git log
- SCP theming with proper incident reports and containment documentation
- Educational progression that matches official Git tutorial structure
- Clean error handling with Foundation-appropriate terminology
- Cross-platform compatibility
