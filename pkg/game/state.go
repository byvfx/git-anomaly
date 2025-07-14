package game

import (
	"time"
)

// GameState represents the entire game state including repository simulation
// and SCP-specific mechanics
type GameState struct {
	// Repository simulation
	IsInitialized bool
	CurrentBranch string
	Branches      map[string][]string // branch -> commit IDs

	// Working directory and staging
	WorkingDir  map[string]FileState
	StagingArea map[string]FileState

	// Commits and history
	Commits     []Commit
	CommitGraph map[string][]string // commit -> parents

	// Git config
	ConfigName  string
	ConfigEmail string

	// SCP-specific state
	AnomalyLevel      int    // 0-100, increases with mistakes
	ContainmentStatus string // "SECURE", "BREACH", "CRITICAL"

	// Progress tracking
	CurrentLevel    int
	CompletedLevels []int
	Score           int
}

// FileState represents the state of a file in the working directory or staging area
type FileState struct {
	Content  string
	Modified bool
	Staged   bool
	Hash     string // Simple hash for change detection
}

// Commit represents a git commit in the simulated repository
type Commit struct {
	ID        string
	Message   string
	Author    string
	Timestamp time.Time
	Files     map[string]string // filename -> hash
	Branch    string
}

// NewGameState creates a new game state with default values
func NewGameState() *GameState {
	return &GameState{
		IsInitialized:     false,
		CurrentBranch:     "",
		Branches:          make(map[string][]string),
		WorkingDir:        make(map[string]FileState),
		StagingArea:       make(map[string]FileState),
		Commits:           []Commit{},
		CommitGraph:       make(map[string][]string),
		AnomalyLevel:      0,
		ContainmentStatus: "SECURE",
		CurrentLevel:      1,
		CompletedLevels:   []int{},
		Score:             0,
	}
}

// UpdateContainmentStatus updates the containment status based on anomaly level
func (gs *GameState) UpdateContainmentStatus() {
	switch {
	case gs.AnomalyLevel >= 80:
		gs.ContainmentStatus = "CRITICAL"
	case gs.AnomalyLevel >= 50:
		gs.ContainmentStatus = "BREACH"
	default:
		gs.ContainmentStatus = "SECURE"
	}
}

// IncreaseAnomaly increases the anomaly level and updates containment status
func (gs *GameState) IncreaseAnomaly(delta int) {
	increase := int(float64(100-gs.AnomalyLevel) * 0.25)
	if increase < delta {
		increase = delta
	}
	gs.AnomalyLevel += increase
	if gs.AnomalyLevel > 100 {
		gs.AnomalyLevel = 100
	}
	gs.UpdateContainmentStatus()
}
