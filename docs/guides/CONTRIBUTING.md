# Contributing to SCP-████: Git Anomaly

Thank you for your interest in contributing to the SCP Foundation's Git learning project!

## Development Setup

### Prerequisites
- Go 1.21 or higher
- golangci-lint (for code quality)

### Building the Project
```bash
# Clone the repository
git clone https://github.com/byvfx/git-anomaly.git
cd git-anomaly

# Build the project
make build

# Run tests
make test
```

## Code Quality

### Linting
We use `golangci-lint` for code quality checks. The project includes a `.golangci.yml` configuration file with v2 format.

```bash
# Run linter
golangci-lint run ./...

# Auto-fix issues where possible
golangci-lint run --fix ./...
```

#### Special Considerations
- **TUI Code**: The Bubble Tea UI code in `pkg/ui/bubbletea.go` is currently on hold for development. Unused functions and variables are commented out but preserved for future development.
- **Terminal UI**: Print statement error checks are disabled for `pkg/ui/terminal.go` as these rarely fail in practice for a CLI game.

### Testing
```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Run with race detection
go test -race ./...
```

## Project Structure

```
scp-git-game/
├── cmd/                    # CLI commands
├── pkg/
│   ├── game/              # Core game logic
│   ├── scp/               # SCP theming and documents
│   └── ui/                # User interface (terminal & TUI)
├── docs/                  # Documentation
├── assets/                # Game assets
└── scripts/               # Build and installation scripts
```

## Development Guidelines

### Code Style
- Follow standard Go conventions
- Use `gofmt` for formatting
- Write meaningful commit messages
- Keep functions focused and testable

### SCP Theming
- Maintain immersive SCP Foundation atmosphere
- Use appropriate terminology (containment, breach, anomaly, etc.)
- Follow SCP documentation format for game text

### Git Learning Focus
- Ensure commands teach real Git concepts
- Follow the official Git tutorial progression
- Provide clear educational feedback

## Submitting Changes

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Run tests and linting
5. Commit with descriptive messages
6. Push to your fork
7. Create a Pull Request

## Issues and Bugs

When reporting issues, please include:
- Go version
- Operating system
- Steps to reproduce
- Expected vs actual behavior
- Any error messages

## Feature Requests

We welcome feature requests! Please:
- Check existing issues first
- Describe the educational value
- Consider the SCP theming
- Provide implementation ideas if possible

## Release Process

The project uses automated releases via GitHub Actions:
- Update `release.md` with release notes
- Create versioned documentation in `docs/releases/`
- Follow semantic versioning (vX.Y.Z)

---

*Secure. Contain. Protect.*