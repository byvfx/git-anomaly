# SCP-████: The Self-Modifying Codebase
## Detailed Implementation Planning

### 🎯 Phase 1 Deep Dive: Foundation Setup

#### 1.1 Core Data Structures

**Game State Management (`pkg/game/state.go`)**
```go
type GameState struct {
    // Repository simulation
    IsInitialized    bool
    CurrentBranch    string
    Branches         map[string][]string // branch -> commit IDs
    
    // Working directory and staging
    WorkingDir       map[string]FileState
    StagingArea      map[string]FileState
    
    // Commits and history
    Commits          []Commit
    CommitGraph      map[string][]string // commit -> parents
    
    // SCP-specific state
    AnomalyLevel     int    // 0-100, increases with mistakes
    ResearcherSanity int    // 0-100, maintained with bubble tea
    ContainmentStatus string // "SECURE", "BREACH", "CRITICAL"
    
    // Progress tracking
    CurrentLevel     int
    CompletedLevels  []int
    Score           int
}

type FileState struct {
    Content   string
    Modified  bool
    Staged    bool
    Hash      string // Simple hash for change detection
}

type Commit struct {
    ID        string
    Message   string
    Author    string
    Timestamp time.Time
    Files     map[string]string // filename -> hash
    Branch    string
}
```

**Level System (`pkg/game/levels.go`)**
```go
type Level struct {
    ID           int
    Title        string
    SCPNumber    string    // e.g., "SCP-████"
    Description  string
    Objective    string
    
    // SCP Documentation
    ObjectClass     string    // Safe, Euclid, Keter
    ContainmentProcs string
    IncidentReport  string
    
    // Game mechanics
    InitialFiles    map[string]string
    RequiredCommands []string
    ValidateFunc    func(*GameState) (bool, string)
    
    // Rewards
    ScoreReward     int
    UnlocksNext     []int
}
```

#### 1.2 SCP Documentation System (`pkg/scp/document.go`)

**Document Templates**
```go
type SCPDocument struct {
    Number      string
    ObjectClass string
    Procedures  string
    Description string
    Addendum    []string
}

func (d *SCPDocument) Format() string {
    return fmt.Sprintf(`
██████╗ ███████╗ ██████╗██████╗       ████████╗██╗  ██╗███████╗
██╔══██╗██╔════╝██╔════╝██╔══██╗      ╚══██╔══╝██║  ██║██╔════╝
███████║█████╗  ██║     ██████╔╝         ██║   ███████║█████╗  
██╔══██║██╔══╝  ██║     ██╔═══╝          ██║   ██╔══██║██╔══╝  
██║  ██║███████╗╚██████╗██║              ██║   ██║  ██║███████╗
╚═╝  ╚═╝╚══════╝ ╚═════╝╚═╝              ╚═╝   ╚═╝  ╚═╝╚══════╝

SCP FOUNDATION SECURE FACILITY - DIGITAL ANOMALIES DIVISION

ITEM #: %s
OBJECT CLASS: %s

SPECIAL CONTAINMENT PROCEDURES: %s

DESCRIPTION: %s
`, d.Number, d.ObjectClass, d.Procedures, d.Description)
}
```

**Incident Report Generator**
```go
func GenerateIncidentReport(level int, error string) string {
    return fmt.Sprintf(`
INCIDENT REPORT - LEVEL %d CONTAINMENT BREACH

RESEARCHER: [REDACTED]
TIMESTAMP: %s
ANOMALY DETECTED: %s

RECOMMENDED ACTION: Review containment procedures and retry operation.
RESEARCHER WELLNESS CHECK: Bubble tea break recommended.

- Dr. ████████, Senior Containment Specialist
`, level, time.Now().Format("2006-01-02 15:04:05"), error)
}
```

#### 1.3 Git Command Simulation (`pkg/game/commands.go`)

**Command Interface**
```go
type GitCommand interface {
    Execute(args []string, state *GameState) CommandResult
    Help() string
    RequiredArgs() int
}

type CommandResult struct {
    Success      bool
    Message      string
    SCPEffect    string  // Special SCP-themed message
    AnomalyDelta int     // Change in anomaly level
    SanityDelta  int     // Change in researcher sanity
}

// Specific command implementations
type InitCommand struct{}
type AddCommand struct{}
type CommitCommand struct{}
type StatusCommand struct{}
type BranchCommand struct{}
```

**Example Command Implementation**
```go
func (c *InitCommand) Execute(args []string, state *GameState) CommandResult {
    if state.IsInitialized {
        return CommandResult{
            Success:   false,
            Message:   "Repository already initialized",
            SCPEffect: "⚠️  CONTAINMENT BREACH: Attempting to re-initialize secure repository",
            AnomalyDelta: 5,
        }
    }
    
    state.IsInitialized = true
    state.CurrentBranch = "main"
    state.Branches["main"] = []string{}
    
    return CommandResult{
        Success:   true,
        Message:   "Initialized empty Git repository",
        SCPEffect: "✅ CONTAINMENT ESTABLISHED: Digital anomaly repository secured",
        SanityDelta: 2,
    }
}
```

#### 1.4 Terminal UI System (`pkg/ui/terminal.go`)

**Color Scheme (SCP Foundation Theme)**
```go
var (
    // SCP Foundation color palette
    SCPRed      = color.New(color.FgRed, color.Bold)
    SCPOrange   = color.New(color.FgYellow, color.Bold)
    SCPGreen    = color.New(color.FgGreen)
    SCPBlue     = color.New(color.FgBlue)
    SCPGray     = color.New(color.FgHiBlack)
    SCPWhite    = color.New(color.FgWhite, color.Bold)
    
    // Status indicators
    StatusSecure   = SCPGreen
    StatusBreach   = SCPOrange
    StatusCritical = SCPRed
)
```

**Display Functions**
```go
func DisplayGameStatus(state *GameState) {
    fmt.Println("═══════════════════════════════════════════════════════════")
    
    // Containment status
    statusColor := getStatusColor(state.ContainmentStatus)
    statusColor.Printf("CONTAINMENT STATUS: %s\n", state.ContainmentStatus)
    
    // Current stats
    fmt.Printf("Branch: %s | Anomaly Level: %d%% | Researcher Sanity: %d%%\n", 
        state.CurrentBranch, state.AnomalyLevel, state.ResearcherSanity)
    
    // Working directory status
    if len(state.StagingArea) > 0 {
        SCPBlue.Println("STAGED FOR CONTAINMENT:")
        for filename := range state.StagingArea {
            fmt.Printf("  📁 %s\n", filename)
        }
    }
    
    fmt.Println("═══════════════════════════════════════════════════════════")
}
```

#### 1.5 Bubble Tea Integration (`pkg/ui/bubbletea.go`)

**Tea Selection System**
```go
type TeaType struct {
    Name        string
    Emoji       string
    SanityBoost int
    FocusBoost  int
    Flavor      string
    Effects     []string
}

var TeaMenu = map[string]TeaType{
    "taro": {
        Name:        "Taro Milk Tea",
        Emoji:       "🟣",
        SanityBoost: 15,
        FocusBoost:  5,
        Flavor:      "Creamy and grounding",
        Effects:     []string{"Reduces anomaly detection sensitivity"},
    },
    "matcha": {
        Name:        "Matcha Latte",
        Emoji:       "🟢",
        SanityBoost: 10,
        FocusBoost:  10,
        Flavor:      "Sharp and clarifying",
        Effects:     []string{"Reveals hidden file changes"},
    },
    "brown-sugar": {
        Name:        "Brown Sugar Bubble Tea",
        Emoji:       "🟤",
        SanityBoost: 20,
        FocusBoost:  0,
        Flavor:      "Sweet and comforting",
        Effects:     []string{"Provides courage for risky merge operations"},
    },
}
```

### 🎮 Phase 1 Level Design

#### Level 1: "Initial Containment"
```go
var Level1 = Level{
    ID:          1,
    Title:       "Initial Containment",
    SCPNumber:   "SCP-████",
    ObjectClass: "Safe",
    Description: "A simple anomalous file has been discovered. Establish initial containment.",
    Objective:   "Initialize repository and commit the anomalous file",
    
    InitialFiles: map[string]string{
        "anomaly.txt": "This file writes itself...",
    },
    
    RequiredCommands: []string{"git init", "git add", "git commit"},
    
    ContainmentProcs: `
All personnel must initialize proper version control before 
handling anomalous digital materials. Standard git protocols 
apply with enhanced monitoring.`,
    
    ValidateFunc: func(state *GameState) (bool, string) {
        if !state.IsInitialized {
            return false, "Repository not initialized"
        }
        if len(state.Commits) == 0 {
            return false, "No commits found - anomaly not contained"
        }
        return true, "✅ SCP-████ successfully contained"
    },
}
```

#### Level 2: "Staging Area Protocols"
```go
var Level2 = Level{
    ID:          2,
    Title:       "Staging Area Protocols",
    SCPNumber:   "SCP-████",
    ObjectClass: "Euclid",
    Description: "Multiple anomalous files detected. Learn proper staging procedures.",
    Objective:   "Stage specific files and commit them separately",
    
    InitialFiles: map[string]string{
        "safe_file.txt":      "Normal file content",
        "anomaly_1.txt":      "I change when you're not looking",
        "anomaly_2.txt":      "I multiply when staged",
        "classified.txt":     "[DATA EXPUNGED]",
    },
    
    RequiredCommands: []string{"git add", "git status", "git commit"},
    
    // More complex validation logic...
}
```

### 🧪 Testing Strategy Details

#### Unit Test Examples
```go
func TestInitCommand(t *testing.T) {
    state := &GameState{}
    cmd := &InitCommand{}
    
    result := cmd.Execute([]string{}, state)
    
    assert.True(t, result.Success)
    assert.True(t, state.IsInitialized)
    assert.Equal(t, "main", state.CurrentBranch)
    assert.Contains(t, result.SCPEffect, "CONTAINMENT ESTABLISHED")
}

func TestAnomalyLevelProgression(t *testing.T) {
    state := &GameState{AnomalyLevel: 50}
    
    // Test that wrong commands increase anomaly level
    cmd := &InitCommand{}
    state.IsInitialized = true // Already initialized
    
    result := cmd.Execute([]string{}, state)
    
    assert.False(t, result.Success)
    assert.Equal(t, 55, state.AnomalyLevel) // Should increase by 5
}
```

#### Integration Test Scenarios
```go
func TestCompleteLevel1(t *testing.T) {
    game := NewGame()
    
    // Simulate user completing level 1
    commands := []string{
        "git init",
        "git add anomaly.txt",
        "git commit -m 'Initial containment of SCP-████'",
    }
    
    for _, cmd := range commands {
        result := game.ProcessCommand(cmd)
        assert.True(t, result.Success, "Command failed: %s", cmd)
    }
    
    // Verify level completion
    completed, msg := game.CurrentLevel.ValidateFunc(game.State)
    assert.True(t, completed, "Level not completed: %s", msg)
}
```

### 🎨 ASCII Art and Visual Elements

#### SCP Foundation Logo (Small)
```
   ███████╗ ██████╗██████╗ 
   ██╔════╝██╔════╝██╔══██╗
   ███████╗██║     ██████╔╝
   ╚════██║██║     ██╔═══╝ 
   ███████║╚██████╗██║     
   ╚══════╝ ╚═════╝╚═╝     
```

#### Containment Status Indicators
```
SECURE:    🟢 ████████████ 100%
BREACH:    🟡 ██████░░░░░░  60%
CRITICAL:  🔴 ███░░░░░░░░░  30%
```

#### Bubble Tea Animation Frames
```
Frame 1: [ ○ ]    Frame 2: [ ● ]    Frame 3: [ ◉ ]
         [   ]             [   ]             [ ○ ]
         [___]             [___]             [___]
```

### 🚀 MVP Implementation Order

1. **Core Game Loop** - Basic CLI that can accept commands
2. **Git Init/Add/Commit** - Essential git commands
3. **SCP Document Formatting** - Basic theming
4. **Level 1 Implementation** - First playable level
5. **Error Handling** - Proper SCP-themed error messages
6. **Basic Testing** - Unit tests for core functions

This should give us a working proof-of-concept that validates the core concept while being achievable in Phase 1.

### 🔧 Development Notes for Claude Code

**Priority Order for Implementation:**
1. Start with `pkg/game/state.go` - Get the data structures right
2. Implement basic commands in `pkg/game/commands.go`
3. Add SCP formatting in `pkg/scp/document.go`
4. Build the CLI loop in `cmd/root.go`
5. Test with Level 1 scenario

**Testing Strategy:**
- Test each component in isolation first
- Use table-driven tests for command validation
- Mock user input for integration tests
- Test error conditions thoroughly (they're part of the SCP experience!)

**Common Gotchas:**
- Ensure terminal colors work across platforms
- Handle edge cases in git command parsing
- Make sure ASCII art renders correctly in different terminals
- Test bubble tea integration doesn't conflict with main CLI